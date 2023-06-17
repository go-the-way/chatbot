package pkg

import "os"

func BindEnv(name, defVal string) (bindVal string) {
	bindVal = defVal
	if val := os.Getenv(name); val != "" {
		bindVal = val
	}
	return
}
