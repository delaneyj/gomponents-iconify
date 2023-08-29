package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Email(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5V4a1 1 0 0 0-1 1h1Zm18 0h1a1 1 0 0 0-1-1v1ZM3 6h18V4H3v2Zm17-1v12h2V5h-2Zm-1 13H5v2h14v-2ZM4 17V5H2v12h2Zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2Zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 5l9 9l9-9"/></g>`),
		g.Group(children),
	)
}