package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M42 6H6a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" d="M4 18h40m-26.5 0v24m13-24v24M4 30h40m0-22v32a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8"/></g>`),
		g.Group(children),
	)
}