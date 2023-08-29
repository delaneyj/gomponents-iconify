package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddLens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.992 6h3m3 0h-3m0 0V3m0 3v3m-3.88 4.5C2.835 18.311 6.987 22 12 22c5.523 0 10-4.477 10-10c0-5.013-3.689-9.165-8.5-9.888"/><path d="M17.197 9c-.1-.172-.207-.34-.323-.5m.937 5a6.01 6.01 0 0 1-4.311 4.311"/></g>`),
		g.Group(children),
	)
}