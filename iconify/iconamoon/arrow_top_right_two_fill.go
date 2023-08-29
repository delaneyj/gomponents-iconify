package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRightTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 16V8a1 1 0 0 0-1-1H8a1 1 0 0 0-.707 1.707l8 8A1 1 0 0 0 17 16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}