package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTwentyThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.838 25a5.931 5.931 0 0 1 5.777-6a5.971 5.971 0 0 1 4.223 10.222c-2.445 2-10 7.778-10 7.778h11.777m10.103-9.111a4.457 4.457 0 0 1 4.444 4.444h0a4.457 4.457 0 0 1-4.444 4.445H30.94c-3.111 0-4.222-.445-5.556-1.556m.001-14.666C26.718 19.445 28.05 19 30.94 19h1.778a4.457 4.457 0 0 1 4.444 4.444h0a4.457 4.457 0 0 1-4.444 4.445h-4.444"/><circle cx="32.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}