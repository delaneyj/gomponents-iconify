package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m2.5 2.5l2 3m8-3l-2 3M4 9.5h1m5 0h1m-9.5 3v-2a6 6 0 1 1 12 0v2h-12Z"/>`),
		g.Group(children),
	)
}