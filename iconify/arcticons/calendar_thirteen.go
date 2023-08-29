package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.94 37h9m-9-15.525L16.44 19m0 0v18m15.176-9.111a4.457 4.457 0 0 1 4.444 4.444h0a4.457 4.457 0 0 1-4.444 4.445h-1.778c-3.11 0-4.222-.445-5.555-1.556m0-14.666C25.616 19.445 26.949 19 29.838 19h1.778a4.457 4.457 0 0 1 4.444 4.444h0a4.457 4.457 0 0 1-4.444 4.445h-4.444"/><circle cx="32.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}