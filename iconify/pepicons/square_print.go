package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquarePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M8 7h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H8a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 8 7Z" opacity=".8"/><path fill-rule="evenodd" d="M14 5.5H6a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5Zm-7.5 8v-7h7v7h-7Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}