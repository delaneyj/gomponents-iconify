package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.444 27.889a4.457 4.457 0 0 1 4.445 4.444h0a4.457 4.457 0 0 1-4.445 4.445h-1.777c-3.112 0-4.223-.445-5.556-1.556m0-14.666C19.444 19.445 20.778 19 23.667 19h1.777a4.457 4.457 0 0 1 4.445 4.444h0a4.457 4.457 0 0 1-4.445 4.445H21"/><circle cx="32.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}