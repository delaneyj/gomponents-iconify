package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonCentered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 15h14v2H9z"/><path fill="currentColor" d="M28 22H4a2.002 2.002 0 0 1-2-2v-8a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v8a2.002 2.002 0 0 1-2 2ZM4 12v8h24v-8Z"/>`),
		g.Group(children),
	)
}