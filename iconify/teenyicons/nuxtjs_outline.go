package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NuxtjsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="m.5 12.5l6-10l6 10H.5Z"/><path d="m4.5 12.5l5-8.5l5 8.5h-10Z"/></g>`),
		g.Group(children),
	)
}