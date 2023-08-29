package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveDownLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.386 3.143a.75.75 0 1 0-.772-1.286c-1.347.808-2.34 1.785-2.98 3.135c-.63 1.33-.884 2.955-.884 5.008v5.44l-2.72-2.72a.75.75 0 1 0-1.06 1.06l4 4a.75.75 0 0 0 1.06 0l4-4a.75.75 0 0 0-1.06-1.06l-2.72 2.72V10c0-1.947.245-3.321.74-4.366c.486-1.026 1.243-1.799 2.396-2.49Z"/>`),
		g.Group(children),
	)
}