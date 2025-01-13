package goversion

import "runtime"

func Get_Go_Version() string {
	return runtime.Version()
}
