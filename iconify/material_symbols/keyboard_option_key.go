package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardOptionKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.775 19L7.85 7H3V5h6l6.925 12H21v2h-6.225ZM15 7V5h6v2h-6Z"/>`),
		g.Group(children),
	)
}