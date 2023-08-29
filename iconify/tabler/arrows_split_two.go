package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsSplitTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 17h-5.397a5 5 0 0 1-4.096-2.133l-.514-.734A5 5 0 0 0 6.897 12H3m18-5h-5.395a5 5 0 0 0-4.098 2.135l-.51.73A5 5 0 0 1 6.9 12H3"/><path d="m18 10l3-3l-3-3m0 16l3-3l-3-3"/></g>`),
		g.Group(children),
	)
}