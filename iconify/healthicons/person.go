package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Person(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M6 36c0-4.965 11.992-8 18-8c6.008 0 18 3.035 18 8v6H6v-6Z"/><path fill-rule="evenodd" d="M24 26c5.523 0 10-4.477 10-10S29.523 6 24 6s-10 4.477-10 10s4.477 10 10 10Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}