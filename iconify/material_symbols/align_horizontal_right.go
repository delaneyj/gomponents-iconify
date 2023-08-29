package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22V2h2v20h-2ZM8 17v-3h10v3H8Zm-6-7V7h16v3H2Z"/>`),
		g.Group(children),
	)
}