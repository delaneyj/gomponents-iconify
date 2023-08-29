package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.414 17.586L18 22.172V8H8V6h10a2.002 2.002 0 0 1 2 2v14.172l4.586-4.586L26 19l-7 7l-7-7Z"/>`),
		g.Group(children),
	)
}