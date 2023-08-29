package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-1H2v1h1Zm0 1h3v-2H3v2Zm3 0v3h2v-3H6Zm-2 3v-4H2v4h2Zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2Zm1-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3H6Zm0-3h2a2 2 0 0 0-2-2v2Zm15-1h1v-1h-1v1Zm-3 1h3v-2h-3v2Zm2-1v4h2v-4h-2Zm-2 4v-3h-2v3h2Zm1 1a1 1 0 0 1-1-1h-2a3 3 0 0 0 3 3v-2Zm1-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm-2-5a2 2 0 0 0-2 2h2v-2Z"/></g>`),
		g.Group(children),
	)
}