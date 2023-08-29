package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRoundedPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.54 20.996C12.364 21 12.184 21 12 21c-7.2 0-9-1.8-9-9s1.8-9 9-9s9 1.8 9 9c0 .185-.001.366-.004.544M16 19h6m-3-3v6"/>`),
		g.Group(children),
	)
}