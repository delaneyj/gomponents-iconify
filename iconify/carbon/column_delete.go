package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 3.41L28.59 2L25 5.59L21.41 2L20 3.41L23.59 7L20 10.59L21.41 12L25 8.41L28.59 12L30 10.59L26.41 7L30 3.41z"/><path fill="currentColor" d="M24 14v14h-6V6h-2v22a2.002 2.002 0 0 0 2 2h6a2.002 2.002 0 0 0 2-2V14zM10 30H4a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v20a2.002 2.002 0 0 1-2 2zM4 8v20h6V8z"/>`),
		g.Group(children),
	)
}