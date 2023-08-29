package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-4.5-4.5l1.45-1.45L12 18.1l3.05-3.05l1.45 1.45L12 21ZM8.95 9.05L7.5 7.6L12 3.1l4.5 4.5l-1.45 1.45L12 6L8.95 9.05Z"/>`),
		g.Group(children),
	)
}