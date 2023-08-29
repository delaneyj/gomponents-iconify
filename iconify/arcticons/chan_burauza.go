package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChanBurauza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.462 24.472c0-2.398.488-4.686 1.6-7.114L4.5 6.156L15.94 9.01C19.148 7.59 21.436 7.1 24 7.1s4.853.489 8.06 1.91L43.5 6.156l-3.562 11.202c1.112 2.428 1.6 4.716 1.6 7.114c0 9.595-7.777 17.372-17.538 17.372S6.462 34.066 6.462 24.472Z"/><circle cx="31.583" cy="23.671" r="3.621" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.985" cy="23.671" r="3.621" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}