package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkCellRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.425 22q-.675 0-.938-.613T3.7 20.3L20.3 3.7q.475-.475 1.088-.213t.612.938V21q0 .425-.288.713T21 22H4.425ZM17 20h3V6.825l-3 3V20Z"/>`),
		g.Group(children),
	)
}