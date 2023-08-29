package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StructureBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M8 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm14 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0ZM8 19a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm14 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path stroke-linecap="round" d="M8 5h8m3 11v-3m0-5v1m-3 10h-3m-3 0H8m-3-3V8"/></g>`),
		g.Group(children),
	)
}