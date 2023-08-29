package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCabinet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M42 4H6v10h36V4Zm0 15H6v10h36V19Zm0 15H6v10h36V34Z"/><path stroke-linecap="round" d="M21 9h6m-6 15h6m-6 15h6"/></g>`),
		g.Group(children),
	)
}