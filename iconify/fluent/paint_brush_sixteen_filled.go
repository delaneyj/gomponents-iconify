package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBrushSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 4.5V1H9v2.5a.5.5 0 0 1-1 0V1H3.5a.5.5 0 0 0-.5.5V7h10V1.5a.5.5 0 0 0-.5-.5H11v3.5a.5.5 0 0 1-1 0ZM13 8H3v1a2 2 0 0 0 2 2h1.5v2.5a1.5 1.5 0 0 0 3 0V11H11a2 2 0 0 0 2-2V8Z"/>`),
		g.Group(children),
	)
}