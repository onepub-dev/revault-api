//go:build windows && arm64

package revault

/*
#cgo LDFLAGS: ${SRCDIR}/native/windows-aarch64-msvc/revault_api.lib -lbcrypt -lcrypt32 -lncrypt -luserenv -lws2_32 -ladvapi32
*/
import "C"
