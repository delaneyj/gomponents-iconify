package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.999 12c0 5.523 4.477 10 10 10s10-4.477 10-10s-4.477-10-10-10s-10 4.477-10 10Zm14.53-.28a.75.75 0 0 1 .073.976l-.072.085l-4.001 4a.75.75 0 0 1-.977.073l-.084-.073l-4-4.001a.75.75 0 0 1 .977-1.133l.084.072l2.72 2.722V7.75a.75.75 0 0 1 .649-.744L11.999 7a.75.75 0 0 1 .743.648l.007.102v6.69l2.72-2.72a.75.75 0 0 1 .977-.073l.084.073Z"/>`),
		g.Group(children),
	)
}