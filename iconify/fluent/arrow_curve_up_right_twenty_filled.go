package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveUpRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.97 7.03a.75.75 0 1 0 1.06-1.06l-4-4a.75.75 0 0 0-1.06 0l-4 4a.75.75 0 0 0 1.06 1.06l2.72-2.72v5.44c0 1.947-.245 3.321-.74 4.366c-.486 1.026-1.243 1.8-2.396 2.49a.75.75 0 1 0 .772 1.287c1.347-.808 2.34-1.785 2.98-3.134c.63-1.33.884-2.956.884-5.009V4.31l2.72 2.72Z"/>`),
		g.Group(children),
	)
}