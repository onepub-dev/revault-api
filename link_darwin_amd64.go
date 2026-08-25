//go:build darwin && amd64

package revault

/*
#cgo LDFLAGS: ${SRCDIR}/native/macos-x86_64/librevault_api.a -framework Security -framework CoreFoundation -framework Foundation -framework IOKit
*/
import "C"
