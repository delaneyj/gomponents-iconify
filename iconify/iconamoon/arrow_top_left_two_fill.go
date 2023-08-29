package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 7H8a1 1 0 0 0-1 1v8a1 1 0 0 0 1.707.707l8-8A1 1 0 0 0 16 7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}