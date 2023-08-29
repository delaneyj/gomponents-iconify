package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.998 4.18l-3.154-2.425a2.01 2.01 0 0 0-2.807.365l-8.4 10.897a2.003 2.003 0 0 0 .365 2.803l3.153 2.425a2.01 2.01 0 0 0 2.807-.365l8.401-10.897a2.003 2.003 0 0 0-.365-2.803zm-8.45 12.287l-.537.681a.8.8 0 0 1-.639.31a.793.793 0 0 1-.485-.164l-3.153-2.425a.792.792 0 0 1-.303-.53a.788.788 0 0 1 .157-.589l.537-.681a.801.801 0 0 1 .64-.311c.124 0 .309.029.485.164l3.154 2.425a.802.802 0 0 1 .144 1.12z"/>`),
		g.Group(children),
	)
}