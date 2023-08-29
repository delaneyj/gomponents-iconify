package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommandBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M8 13v3h8V8H8v1"/><path d="M16 16.001h3a3 3 0 1 1-3 3v-3Zm-7.999 0h-3a3 3 0 1 0 3 3v-3ZM16 8h3a3 3 0 1 0-3-3v3ZM8.001 8h-3a3 3 0 1 1 3-3v3Z"/></g>`),
		g.Group(children),
	)
}