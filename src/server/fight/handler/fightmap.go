package handler

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	. "gserver_service_demo/src/server/fight/cfg"
	. "gserver_service_demo/src/server/fight/entity"

	. "github.com/gfandada/gserver/gameutil/entity"
	//	"github.com/gfandada/gserver/logger"
)

const (
	RED_SHANG  = 31 // 红上
	RED_ZHONG  = 32 // 红中
	RED_XIA    = 33 // 红下
	BLUE_SHANG = 34 // 红上
	BLUE_ZHONG = 35 // 红中
	BLUE_XIA   = 36 // 红下
)

type FightMap struct {
	index  map[EntityId]int    // 当前索引
	points map[int][]*WayPoint // 点序
}

func (f *FightMap) Load() {
	f.index = make(map[EntityId]int)
	f.points = make(map[int][]*WayPoint)
	world := ParseWorldByCSV("server/fight/handler/地图数据-路径点.csv")
	points1 := world.AllOfKind(RED_SHANG)
	sort.Sort(REDWayPoint(points1))
	points2 := world.AllOfKind(RED_ZHONG)
	sort.Sort(REDWayPoint(points2))
	points3 := world.AllOfKind(RED_XIA)
	sort.Sort(REDWayPoint(points3))
	points4 := world.AllOfKind(BLUE_SHANG)
	sort.Sort(BLUEWayPoint(points4))
	points5 := world.AllOfKind(BLUE_ZHONG)
	sort.Sort(BLUEWayPoint(points5))
	points6 := world.AllOfKind(BLUE_XIA)
	sort.Sort(BLUEWayPoint(points6))
	for i := RED_SHANG; i <= BLUE_XIA; i++ {
		points := world.AllOfKind(i)
		for _, v := range points {
			v.Kind = KindPlain
		}
	}
	//	f.setPath(RED_SHANG, world, points1)
	//	f.setPath(RED_ZHONG, world, points2)
	//	f.setPath(RED_XIA, world, points3)
	//	f.setPath(BLUE_SHANG, world, points4)
	//	f.setPath(BLUE_ZHONG, world, points5)
	//	f.setPath(BLUE_XIA, world, points6)
	f.setPath(RED_SHANG, world, points4)
	f.setPath(RED_ZHONG, world, points5)
	f.setPath(RED_XIA, world, points6)
	f.setPath(BLUE_SHANG, world, points1)
	f.setPath(BLUE_ZHONG, world, points2)
	f.setPath(BLUE_XIA, world, points3)
}

func (f *FightMap) getPoint(flag, index int, entityId EntityId) (int, int) {
	f.index[entityId] = index
	point := f.points[flag][index]
	return point.X, point.Y
}

func (f *FightMap) setPath(flag int, world World, points []*WayPoint) {
	var newPoints []*WayPoint
	var newPoints1 []Pather
	for i, v := range points {
		if i != len(points)-1 {
			p, _, found := Path(v, points[i+1])
			if found {
				newPoints = append(newPoints, v)
				newPoints1 = append(newPoints1, v)
				for i, _ := range p {
					if i == 0 || i == len(p)-1 {
						continue
					}
					newPoints = append(newPoints, p[len(p)-i-1].(*WayPoint))
					newPoints1 = append(newPoints1, p[len(p)-i-1])
				}
			} else {
				return
			}
		} else {
			newPoints = append(newPoints, v)
			newPoints1 = append(newPoints1, v)
		}
	}
	f.points[flag] = newPoints
	//f.printPoints(newPoints)
	file, err := os.OpenFile("path"+strconv.Itoa(flag)+".csv", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		file.WriteString("")
		file.WriteString(world.RenderPath(newPoints1))
	}
}

func (f *FightMap) printPoints(points []*WayPoint) {
	for _, v := range points {
		fmt.Print("[", v.X, ":", v.Y, "] ")
	}
	fmt.Print("\n")
}

func (f *FightMap) Unload() {
	f.index = nil
	f.points = nil
}

func (f *FightMap) haveEnemy(soldier *Entity) bool {
	neighbors := soldier.Neighbors()
	_, radius, _ := GetSoldierAttack(uint32(1))
	if AttackJudge(soldier, neighbors, int(radius)) == nil {
		return false
	}
	return true
}

// FIXME 取巧地使用了pos.Y来标识两个阵营的不同路线小兵
func (f *FightMap) updateIndex(soldiers map[*Entity]struct{}, time int64) {
	var newpos Vector3
	for soldier := range soldiers {
		pos := soldier.GetPosition()
		if pos.TIME == 0 {
			pos.TIME = time
		}
		if !f.haveEnemy(soldier) {
			// 获取当前点
			currPoint := f.points[int(pos.Y)][f.index[soldier.Id]]
			if f.index[soldier.Id]+1 >= len(f.points[int(pos.Y)]) {
				continue
			}
			// 获取下一个点
			nextPoint := f.points[int(pos.Y)][f.index[soldier.Id]+1]
			timej := Coord(time-pos.TIME) / Coord(1e3)
			// X轴
			if currPoint.X == nextPoint.X {
				// Y轴正方向
				if nextPoint.Y > currPoint.Y {
					newpos = Vector3{
						X:    Coord(pos.X),
						Y:    pos.Y,
						Z:    Coord(pos.Z + Coord(1)*timej),
						TIME: time,
					}
					if float32(nextPoint.Y)-float32(newpos.Z) < 0.1 {
						f.index[soldier.Id] += 1
					}
					// Y轴负方向
				} else {
					newpos = Vector3{
						X:    Coord(pos.X),
						Y:    pos.Y,
						Z:    Coord(pos.Z + Coord(-1)*timej),
						TIME: time,
					}
					if float32(newpos.Z)-float32(nextPoint.Y) < 0.1 {
						f.index[soldier.Id] += 1
					}
				}
				// Y轴
			} else if currPoint.Y == nextPoint.Y {
				// X轴正方向
				if nextPoint.X > currPoint.X {
					newpos = Vector3{
						X:    Coord(pos.X + Coord(1)*timej),
						Y:    pos.Y,
						Z:    Coord(pos.Z),
						TIME: time,
					}
					if float32(nextPoint.X)-float32(newpos.X) < 0.1 {
						f.index[soldier.Id] += 1
					}
					// X轴负方向
				} else {
					newpos = Vector3{
						X:    Coord(pos.X + Coord(-1)*timej),
						Y:    pos.Y,
						Z:    Coord(pos.Z),
						TIME: time,
					}
					if float32(newpos.X)-float32(nextPoint.X) < 0.1 {
						f.index[soldier.Id] += 1
					}
				}
			}
			soldier.MoveSpace(newpos)
			//			str := fmt.Sprintf("小兵当前位置{%v}新位置{%v},规划的点序{%d,%d}=>{%d,%d}",
			//				pos, newpos, currPoint.X, currPoint.Y, nextPoint.X, nextPoint.Y)
			//			logger.Info(str)
		} else {
			newpos := Vector3{
				X:    Coord(pos.X),
				Y:    pos.Y,
				Z:    Coord(pos.Z),
				TIME: time,
			}
			soldier.MoveSpace(newpos)
		}
	}
}

type REDWayPoint []*WayPoint

func (c REDWayPoint) Len() int {
	return len(c)
}
func (c REDWayPoint) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c REDWayPoint) Less(i, j int) bool {
	return c[i].X+c[i].Y < c[j].X+c[j].Y
}

type BLUEWayPoint []*WayPoint

func (c BLUEWayPoint) Len() int {
	return len(c)
}
func (c BLUEWayPoint) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c BLUEWayPoint) Less(i, j int) bool {
	return c[i].X+c[i].Y > c[j].X+c[j].Y
}
