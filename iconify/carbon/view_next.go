package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.587 22L15 16.41V7h1.998v8.582L22 20.587L20.587 22z"/><path fill="currentColor" d="M16 2a13.916 13.916 0 0 1 10 4.234V2h2v8h-8V8h4.922A11.982 11.982 0 1 0 28 16h2A14 14 0 1 1 16 2Z"/>`),
		g.Group(children),
	)
}