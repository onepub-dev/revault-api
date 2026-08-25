//go:build darwin && arm64

package revault

/*
#cgo LDFLAGS: ${SRCDIR}/native/macos-aarch64/librevault_api.a -framework Security -framework CoreFoundation -framework Foundation -framework IOKit
*/
import "C"
