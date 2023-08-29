package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallOfDuty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 10.467v27.01l4.116 2.873l6.296-4.034l.003-25.145L8.7 7.651Zm39 0v27.01l-4.116 2.873l-6.296-4.034l-.003-25.145l6.215-3.52Zm-13.317 2.301v14.575L24 30.846l-6.184-3.503V12.768L24 16.27Z"/>`),
		g.Group(children),
	)
}