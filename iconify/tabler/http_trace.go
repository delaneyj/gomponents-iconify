package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HttpTrace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8h4M5 8v8m5-4h2a2 2 0 1 0 0-4h-2v8m4 0l-3-4m6 4v-6a2 2 0 1 1 4 0v6m-4-3h4"/>`),
		g.Group(children),
	)
}