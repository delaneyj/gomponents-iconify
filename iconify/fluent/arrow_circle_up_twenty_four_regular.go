package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m7.47 12.28l.084.073a.75.75 0 0 0 .977-.073l2.72-2.72v6.69l.007.102A.75.75 0 0 0 12 17l.101-.006a.75.75 0 0 0 .649-.744V9.56l2.72 2.722l.084.072a.75.75 0 0 0 .977-1.133l-4-4.001l-.084-.073a.75.75 0 0 0-.977.073l-4 4l-.073.085a.75.75 0 0 0 .072.976ZM22.001 12c0-5.523-4.477-10-10-10s-10 4.477-10 10s4.477 10 10 10s10-4.477 10-10Zm-18.5 0a8.5 8.5 0 1 1 17 0a8.5 8.5 0 0 1-17 0Z"/>`),
		g.Group(children),
	)
}