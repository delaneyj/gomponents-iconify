package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StructureLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M8 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm14 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0ZM8 19a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm14 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path stroke-linecap="round" d="M8 19h8M8 5h8m3 11V8M5 16V8" opacity=".5"/></g>`),
		g.Group(children),
	)
}