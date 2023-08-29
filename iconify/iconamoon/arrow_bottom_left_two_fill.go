package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeftTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 8v8a1 1 0 0 0 1 1h8a1 1 0 0 0 .707-1.707l-8-8A1 1 0 0 0 7 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}