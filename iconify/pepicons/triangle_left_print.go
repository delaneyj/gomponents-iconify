package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="m7.735 10.576l8-5A.5.5 0 0 1 16.5 6v10a.5.5 0 0 1-.765.424l-8-5a.5.5 0 0 1 0-.848Z" opacity=".8"/><path fill-rule="evenodd" d="m13.735 4.576l-8 5a.5.5 0 0 0 0 .848l8 5A.5.5 0 0 0 14.5 15V5a.5.5 0 0 0-.765-.424ZM13.5 5.902v8.196L6.943 10L13.5 5.902Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}