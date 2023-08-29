package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 36.526H21c1.703 0 2.345 1.498 3 2.55c.645-1.024 1.303-2.55 3-2.55h15.5m-2.846-8.202c0 2.613-3.024 5.011-6.41 5.011a5.093 5.093 0 0 1-2.385-.371l-1.749.822l.164-1.743c-1.45-.725-1.92-2.786-1.88-3.719c.113-2.611 2.745-4.732 6.13-4.732s6.13 2.119 6.13 4.732Z"/><circle cx="32.303" cy="27.575" r=".75" fill="currentColor"/><circle cx="35.521" cy="27.575" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}