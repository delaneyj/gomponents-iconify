package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M5 9.5V19a3 3 0 0 1-.6 1.8L3.5 22M19 9.5V19a3 3 0 0 0 .6 1.8l.9 1.2"/><path d="M18 9.5c0 4.142-2.686 7.5-6 7.5s-6-3.358-6-7.5C6 5.358 8.686 2 12 2s6 3.358 6 7.5ZM5 20h14"/><path stroke-linecap="round" d="M13 5.256c.96.51 1.697 1.732 1.926 3.244"/></g>`),
		g.Group(children),
	)
}