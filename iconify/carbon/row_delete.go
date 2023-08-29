package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 30H4a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2zM4 22h-.002L4 28h20v-6zM30 3.41L28.59 2L25 5.59L21.41 2L20 3.41L23.59 7L20 10.59L21.41 12L25 8.41L28.59 12L30 10.59L26.41 7L30 3.41z"/><path fill="currentColor" d="M4 14V8h14V6H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h22v-2Z"/>`),
		g.Group(children),
	)
}