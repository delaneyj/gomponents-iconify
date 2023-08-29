package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 13h.01M16 13h.01M13 14.5h.01M10 16h.01M16 16h.01M13 17.5h.01M8 7.036a3.484 3.484 0 0 1 1.975.99M6 13.662A3.5 3.5 0 0 1 8.37 7.11a5.002 5.002 0 0 1 9.614 1.49a2.751 2.751 0 0 1 1.59 4.122"/>`),
		g.Group(children),
	)
}