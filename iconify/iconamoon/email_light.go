package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5v-.75a.75.75 0 0 0-.75.75H3Zm18 0h.75a.75.75 0 0 0-.75-.75V5ZM3 5.75h18v-1.5H3v1.5ZM20.25 5v12h1.5V5h-1.5ZM19 18.25H5v1.5h14v-1.5ZM3.75 17V5h-1.5v12h1.5ZM5 18.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 5 19.75v-1.5ZM20.25 17c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 21.75 17h-1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 5l9 9l9-9"/></g>`),
		g.Group(children),
	)
}