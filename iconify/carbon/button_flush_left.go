package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonFlushLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 15h14v2H6z"/><path fill="currentColor" d="M28 22H4c-1.103 0-2-.897-2-2v-8c0-1.103.897-2 2-2h24c1.103 0 2 .897 2 2v8c0 1.103-.897 2-2 2zM4 12v8h24v-8H4z"/>`),
		g.Group(children),
	)
}