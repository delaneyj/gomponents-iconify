package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartCombo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 28V16h-8v12h-4V12H7v16H4V2H2v26a2 2 0 0 0 2 2h26v-2Zm-14 0H9V14h4Zm12 0h-4V18h4Z"/><path fill="currentColor" d="M22.786 14a1.988 1.988 0 0 1-1.18-.387L11.205 5.987L8.242 10L6.637 8.806l2.982-4a1.998 1.998 0 0 1 2.749-.446L22.789 12l3.604-4.86L28 8.33l-3.604 4.862a2.001 2.001 0 0 1-1.61.808Z"/>`),
		g.Group(children),
	)
}