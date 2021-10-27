package core

import (
	"fmt"
	"strconv"
)

// Obtiene el puerto si existe a partir de
// los argumentos de entrada del programa
func GetPortArg(args []string, p int) (string, bool) {
	if args == nil {
		return strconv.Itoa(p), false
	}
	for i, arg := range args {
		if arg == "--port" || arg == "-p" {
			if (i + 1) > (len(args) - 1) {
				return strconv.Itoa(p), false
			} else {
				pInt, err := strconv.ParseInt(args[i+1], 10, 32)
				if err != nil {
					return strconv.Itoa(p), false
				} else {
					return strconv.Itoa(int(pInt)), false
				}
			}
		}
	}
	return strconv.Itoa(p), false
}

// Obtiene el prefijo de url de servicios si existe a partir de
// los argumentos de entrada del programa
func GetPrefixArg(args []string, prefix string) (string, bool) {
	if args == nil {
		return prefix, false
	}
	for i, arg := range args {
		if arg == "--prefix" || arg == "-x" {
			if (i + 1) > (len(args) - 1) {
				return prefix, false
			} else {
				return fmt.Sprintf(`%s\/`, args[i+1]), true
			}
		}
	}
	return prefix, false
}
