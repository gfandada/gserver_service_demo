package entity

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/magicsea/behavior3go/config"
)

var (
	_aiManager *AiManager
)

type AiManager struct {
	cfg map[string]*BTTreeCfg
}

func (ai *AiManager) put(key string, value *BTTreeCfg) {
	ai.cfg[key] = value
}

func (ai *AiManager) get(key string) *BTTreeCfg {
	return ai.cfg[key]
}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil || f.IsDir() {
			return nil
		}
		cfg, ok := LoadTreeCfg(path)
		if !ok {
			return fmt.Errorf("load %s failed", path)
		}
		_aiManager.put(f.Name(), cfg)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() %s returned %v\n", path, err)
	}
}

// @params path:ai配置目录
func NewAiManager(dir string) {
	if _aiManager != nil {
		return
	}
	_aiManager = &AiManager{
		cfg: make(map[string]*BTTreeCfg),
	}
	getFilelist(dir)
}
