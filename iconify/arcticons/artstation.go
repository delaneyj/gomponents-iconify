package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Artstation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.066 11.697l8.307 14.762H8.549l8.517-14.762zm25.151 21.099l-3.795 6.575L19.989 6.63h7.126c.757 0 1.456.403 1.834 1.058L42.217 30.68a2.117 2.117 0 0 1 0 2.116Zm-9.568 6.575h-21.52a2.117 2.117 0 0 1-1.834-1.058L5.5 31.738h22.848l4.3 7.634Z"/>`),
		g.Group(children),
	)
}