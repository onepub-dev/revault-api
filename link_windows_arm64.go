//go:build windows && arm64

package revault

/*
#cgo LDFLAGS: ${SRCDIR}/native/windows-aarch64-msvc/librevault_api_go.a -lbcrypt -lcrypt32 -lncrypt -luserenv -lws2_32 -ladvapi32 -lntdll -Wl,-Bstatic -lunwind -Wl,-Bdynamic
*/
import "C"
