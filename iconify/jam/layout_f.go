package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20V8h9v9a3 3 0 0 1-3 3H9zm-2 0H3a3 3 0 0 1-3-3V8h7v12zM18 6H0V3a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v3z"/>`),
		g.Group(children),
	)
}