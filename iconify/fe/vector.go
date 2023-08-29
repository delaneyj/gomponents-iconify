package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feVector0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feVector1" fill="currentColor" fill-rule="nonzero"><path id="feVector2" d="M8 20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2V8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2h8a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2v8a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2H8ZM8 6a2 2 0 0 1-2 2v8a2 2 0 0 1 2 2h8a2 2 0 0 1 2-2V8a2 2 0 0 1-2-2H8ZM4 4v2h2V4H4Zm14 0v2h2V4h-2Zm0 14v2h2v-2h-2ZM4 18v2h2v-2H4Z"/></g></g>`),
		g.Group(children),
	)
}