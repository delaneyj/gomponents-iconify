package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardControlKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.4 13.4L5 12l7-7l7 7l-1.4 1.4L12 7.825L6.4 13.4Z"/>`),
		g.Group(children),
	)
}