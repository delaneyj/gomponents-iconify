package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.385 31a5.975 5.975 0 0 0 6 6a5.79 5.79 0 0 0 5.777-6v-6a5.931 5.931 0 0 0-5.777-6a6.12 6.12 0 0 0-6 6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".997" d="M35.829 21.222L26.94 34.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.17 27.889a4.457 4.457 0 0 1 4.445 4.444h0a4.457 4.457 0 0 1-4.444 4.445h-1.778c-3.111 0-4.222-.445-5.556-1.556m.001-14.666C12.17 19.445 13.504 19 16.393 19h1.778a4.457 4.457 0 0 1 4.444 4.444h0a4.457 4.457 0 0 1-4.444 4.445h-4.445"/><circle cx="32.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}