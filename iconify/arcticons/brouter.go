package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brouter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33 9.5l-1.45 2l2.45 1l-3 1c-.95.32-4.48 9-3 12c1 2 5.17 3.89-10.83 7.89c.39 1.11.83 2.11.83 2.11a6.93 6.93 0 0 1-4 3"/>`),
		g.Group(children),
	)
}