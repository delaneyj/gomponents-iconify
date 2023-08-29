package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="m17.265 11.424l-8 5A.5.5 0 0 1 8.5 16V6a.5.5 0 0 1 .765-.424l8 5a.5.5 0 0 1 0 .848Z" opacity=".8"/><path fill-rule="evenodd" d="m7.265 15.424l8-5a.5.5 0 0 0 0-.848l-8-5A.5.5 0 0 0 6.5 5v10a.5.5 0 0 0 .765.424Zm.235-1.326V5.902L14.057 10L7.5 14.098Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}