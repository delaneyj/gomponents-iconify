package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlameAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 3c3.667 3.333 5.5 6.429 5.5 9.286c0 3.078-1.27 5.198-3.111 6.148c.23-.491.361-1.092.361-1.791c0-1.429-.917-2.976-2.75-4.643c-1.833 1.667-2.75 3.214-2.75 4.643c0 .7.131 1.3.36 1.791c-1.84-.95-3.11-3.07-3.11-6.148C5 9.429 6.833 6.333 10.5 3z"/>`),
		g.Group(children),
	)
}