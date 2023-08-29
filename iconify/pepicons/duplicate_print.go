package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DuplicatePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M15 9h2a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-2H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v2Z" opacity=".8"/><path fill-rule="evenodd" d="M15 7.5H8a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5Zm-6.5 7v-6h6v6h-6Z" clip-rule="evenodd"/><path d="M5.5 11.5h3v1H5a.5.5 0 0 1-.5-.5V5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5v3.5h-1v-3h-6v6Z"/></g>`),
		g.Group(children),
	)
}