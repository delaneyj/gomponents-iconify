package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grocy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.906 18.76H10.094a3 3 0 0 0-2.856 3.92l3.174 9.857a7 7 0 0 0 6.663 4.854h13.85a7 7 0 0 0 6.663-4.854l3.174-9.857a3 3 0 0 0-2.856-3.92Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.868 18.76v-2.019a6.132 6.132 0 0 1 12.264 0v2.019m-2.742 7.645a3.393 3.393 0 0 0-3.572-3.393a3.524 3.524 0 0 0-3.208 3.585v3.148A3.394 3.394 0 0 0 24 33.143h0a3.394 3.394 0 0 0 3.39-3.398H24"/>`),
		g.Group(children),
	)
}