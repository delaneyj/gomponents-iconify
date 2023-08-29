package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBrushOutline0"><g id="evaBrushOutline1"><path id="evaBrushOutline2" fill="currentColor" d="M20 6.83a2.76 2.76 0 0 0-.82-2a2.89 2.89 0 0 0-4 0l-6.6 6.6h-.22a4.42 4.42 0 0 0-4.3 4.31L4 19a1 1 0 0 0 .29.73A1.05 1.05 0 0 0 5 20l3.26-.06a4.42 4.42 0 0 0 4.31-4.3v-.23l6.61-6.6A2.74 2.74 0 0 0 20 6.83ZM8.25 17.94L6 18v-2.23a2.4 2.4 0 0 1 2.4-2.36a2.15 2.15 0 0 1 2.15 2.19a2.4 2.4 0 0 1-2.3 2.34Zm9.52-10.55l-5.87 5.87a4.55 4.55 0 0 0-.52-.64a3.94 3.94 0 0 0-.64-.52l5.87-5.86a.84.84 0 0 1 1.16 0a.81.81 0 0 1 .23.59a.79.79 0 0 1-.23.56Z"/></g></g>`),
		g.Group(children),
	)
}