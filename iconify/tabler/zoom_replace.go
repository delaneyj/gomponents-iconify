package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomReplace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 21l-6-6M3.291 8a7 7 0 0 1 5.077-4.806a7.021 7.021 0 0 1 8.242 4.403"/><path d="M17 4v4h-4m3.705 4a7 7 0 0 1-5.074 4.798a7.021 7.021 0 0 1-8.241-4.403"/><path d="M3 16v-4h4"/></g>`),
		g.Group(children),
	)
}