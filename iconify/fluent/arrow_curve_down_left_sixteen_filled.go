package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveDownLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.151 1.378a.75.75 0 0 1-.279 1.023C8.458 3.781 8.25 6.03 8.25 8.001v4.438l2.22-2.22a.75.75 0 1 1 1.06 1.061l-3.5 3.5a.75.75 0 0 1-1.06 0l-3.5-3.5a.75.75 0 0 1 1.06-1.06l2.22 2.22V8c0-2.029.192-5.08 3.378-6.901a.75.75 0 0 1 1.023.279Z"/>`),
		g.Group(children),
	)
}