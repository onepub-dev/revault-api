//go:build windows && amd64

package revault

/*
#cgo LDFLAGS: ${SRCDIR}/native/windows-x86_64-msvc/librevault_api_go.a -lbcrypt -lcrypt32 -lncrypt -luserenv -lws2_32 -ladvapi32 -lntdll -static-libgcc
*/
import "C"
