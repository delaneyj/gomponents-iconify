package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 2h-4C7.477 2 3 6.477 3 12s4.477 10 10 10h4v-5h-4a5 5 0 0 1 0-10h4V2Z"/><path d="M17 7h2.5A1.5 1.5 0 0 0 21 5.5v-2A1.5 1.5 0 0 0 19.5 2H17v5Zm0 10v5h2.5a1.5 1.5 0 0 0 1.5-1.5v-2a1.5 1.5 0 0 0-1.5-1.5H17Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}