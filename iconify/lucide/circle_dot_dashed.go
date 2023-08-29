package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleDotDashed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.1 2.18a9.93 9.93 0 0 1 3.8 0m3.7 1.53a9.95 9.95 0 0 1 2.69 2.7m1.53 3.69a9.93 9.93 0 0 1 0 3.8m-1.53 3.7a9.95 9.95 0 0 1-2.7 2.69m-3.69 1.53a9.94 9.94 0 0 1-3.8 0m-3.7-1.53a9.95 9.95 0 0 1-2.69-2.7M2.18 13.9a9.93 9.93 0 0 1 0-3.8m1.53-3.7a9.95 9.95 0 0 1 2.7-2.69"/><circle cx="12" cy="12" r="1"/></g>`),
		g.Group(children),
	)
}