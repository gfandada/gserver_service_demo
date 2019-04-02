package cfg

import (
	"fmt"
	"testing"

	"github.com/gfandada/gserver"
)

func Test_cfg(t *testing.T) {
	gserver.RunXlsxService("../../../../cfg/xlsxlogger.xml", "../../../../cfg/xlsx/")
	fmt.Println(GetTowerDura(uint32(1)))
	fmt.Println(GetTowerAttack(uint32(1)))
	fmt.Println(GetTowerDura(uint32(21)))
	fmt.Println(GetTowerAttack(uint32(21)))

	fmt.Println(GetDefaultShip(uint32(1)))
	fmt.Println(GetDefaultShipEquip(uint32(1)))

	fmt.Println(GetEquipDura(uint32(1)))
	fmt.Println(GetEquipAttack(uint32(1)))
}
