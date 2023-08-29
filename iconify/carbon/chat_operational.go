package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatOperational(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.586 10.414L19.166 14l-3.582 3.587L17 19l5-5l-5-5l-1.414 1.414z"/><path fill="currentColor" d="M20.586 10.414L24.166 14l-3.582 3.587L22 19l5-5l-5-5l-1.414 1.414zM10 9l1.593 3l3.407.414l-2.5 2.253L13 18l-3-1.875L7 18l.5-3.333L5 12.414L8.5 12L10 9z"/><path fill="currentColor" d="M17.736 30L16 29l4-7h6a1.997 1.997 0 0 0 2-2V8a1.997 1.997 0 0 0-2-2H6a1.997 1.997 0 0 0-2 2v12a1.997 1.997 0 0 0 2 2h9v2H6a4 4 0 0 1-4-4V8a3.999 3.999 0 0 1 4-4h20a3.999 3.999 0 0 1 4 4v12a4 4 0 0 1-4 4h-4.835Z"/>`),
		g.Group(children),
	)
}