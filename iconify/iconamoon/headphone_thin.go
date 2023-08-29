package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-.5h-.5v.5H3Zm0 .5h3v-1H3v1Zm3.5.5v3h1v-3h-1Zm-3 3v-4h-1v4h1ZM5 20.5A1.5 1.5 0 0 1 3.5 19h-1A2.5 2.5 0 0 0 5 21.5v-1ZM6.5 19A1.5 1.5 0 0 1 5 20.5v1A2.5 2.5 0 0 0 7.5 19h-1ZM6 15.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 6 14.5v1Zm15-.5h.5v-.5H21v.5Zm-3 .5h3v-1h-3v1Zm2.5-.5v4h1v-4h-1Zm-3 4v-3h-1v3h1Zm1.5 1.5a1.5 1.5 0 0 1-1.5-1.5h-1a2.5 2.5 0 0 0 2.5 2.5v-1Zm1.5-1.5a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5h-1ZM18 14.5a1.5 1.5 0 0 0-1.5 1.5h1a.5.5 0 0 1 .5-.5v-1Z"/></g>`),
		g.Group(children),
	)
}