package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.177 13.702l14.771 7.333c1.376.683 2.656 1.792 2.656 3.171c0 1.378-1.28 2.488-2.655 3.17L17.178 34.71"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24c0 11.874-9.626 21.5-21.5 21.5S2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24Z"/>`),
		g.Group(children),
	)
}