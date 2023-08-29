package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5V3.75c-.69 0-1.25.56-1.25 1.25H3Zm18 0h1.25c0-.69-.56-1.25-1.25-1.25V5ZM3 6.25h18v-2.5H3v2.5ZM19.75 5v12h2.5V5h-2.5ZM19 17.75H5v2.5h14v-2.5ZM4.25 17V5h-2.5v12h2.5Zm.75.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 5 20.25v-2.5ZM19.75 17a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 22.25 17h-2.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m3 5l9 9l9-9"/></g>`),
		g.Group(children),
	)
}