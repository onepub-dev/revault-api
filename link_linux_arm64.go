//go:build linux && arm64

package revault

/*
#cgo LDFLAGS: ${SRCDIR}/native/linux-aarch64-gnu/librevault_api.a -ldbus-1 -lpthread -ldl -lm
*/
import "C"
