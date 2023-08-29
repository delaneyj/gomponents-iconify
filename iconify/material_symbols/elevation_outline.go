package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.05 21l6.45-9h5.05L21 3.3V21H2.05Zm1.75-5.825l-1.6-1.15L6.5 8h5.05l4.7-5.475l1.5 1.3L12.45 10H7.5l-3.7 5.175ZM5.95 19H19V8.7L14.45 14H9.5l-3.55 5ZM19 19Z"/>`),
		g.Group(children),
	)
}