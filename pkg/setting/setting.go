package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(paths ...string) (*Setting, error) {
	v := viper.New()
	v.SetConfigName("config")
	for _, path := range paths {
		if path != "" {
			v.AddConfigPath(path)
		}
	}
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{
		vp: v,
	}
	s.WatchConfig()
	return s, nil
}

func (s *Setting) WatchConfig() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(e fsnotify.Event) {
			s.ReloadSection()
		})
	}()
}
