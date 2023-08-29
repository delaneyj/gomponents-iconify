package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m15.194 59.995l69.732-.074v-.014a2.493 2.493 0 0 0 2.361-2.489a2.487 2.487 0 0 0-.802-1.823L51.834 21.02l-.004.004a2.484 2.484 0 0 0-1.902-.892a2.494 2.494 0 0 0-2.02 1.041l-34.46 34.535a2.498 2.498 0 0 0 1.746 4.287zm72.114 17.258l-.01-9.803v-.05h-.005a2.534 2.534 0 0 0-2.534-2.485v-.006l-69.751.074v.042a2.53 2.53 0 0 0-2.293 2.516c0 .033.008.063.01.096l.01 9.477c-.006.074-.022.145-.022.22a2.528 2.528 0 0 0 2.311 2.511v.023l69.751-.074a2.536 2.536 0 0 0 2.534-2.539l-.001-.002z"/>`),
		g.Group(children),
	)
}