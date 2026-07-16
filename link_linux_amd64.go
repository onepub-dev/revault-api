//go:build linux && amd64

package revault

/*
#cgo LDFLAGS: ${SRCDIR}/native/linux-x86_64-gnu/librevault_api.a -ldbus-1 -lpthread -ldl -lm
*/
import "C"
