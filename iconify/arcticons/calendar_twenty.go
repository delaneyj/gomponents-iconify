package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.385 31a5.975 5.975 0 0 0 6 6a5.79 5.79 0 0 0 5.777-6v-6a5.931 5.931 0 0 0-5.777-6a6.12 6.12 0 0 0-6 6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".997" d="M35.829 21.222L26.94 34.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.838 25a5.931 5.931 0 0 1 5.777-6a5.971 5.971 0 0 1 4.223 10.222c-2.445 2-10 7.778-10 7.778h11.777"/><circle cx="32.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}