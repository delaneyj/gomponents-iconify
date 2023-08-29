package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tilde(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12c0-1.657 1.592-3 3.556-3c1.963 0 3.11 1.5 4.444 3c1.333 1.5 2.48 3 4.444 3S20 13.657 20 12"/>`),
		g.Group(children),
	)
}