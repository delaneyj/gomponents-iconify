package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="m12.076 16.265l-5-8A.5.5 0 0 1 7.5 7.5h10a.5.5 0 0 1 .424.765l-5 8a.5.5 0 0 1-.848 0Z" opacity=".8"/><path fill-rule="evenodd" d="m4.576 7.265l5 8a.5.5 0 0 0 .848 0l5-8A.5.5 0 0 0 15 6.5H5a.5.5 0 0 0-.424.765Zm9.522.235L10 14.057L5.902 7.5h8.196Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}