package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aura(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.824 7.594c2.827 2.508 13.391 25.45 13.925 27.317c.533 1.867.373 5.495.373 5.495H42.5L27.881 7.594H15.824Zm2.054 32.812H5.5l5.894-13.23c1.109 3.173 4.777 11.163 6.484 13.23Z"/>`),
		g.Group(children),
	)
}