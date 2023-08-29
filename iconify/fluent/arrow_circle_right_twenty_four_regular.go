package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.001 2c5.523 0 10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10Zm0 1.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17Zm-.353 4.053l.073-.084a.75.75 0 0 1 .976-.073l.084.073l4.001 4a.75.75 0 0 1 .073.977l-.073.085l-4.002 4a.75.75 0 0 1-1.133-.977l.073-.084l2.722-2.721H7.75a.75.75 0 0 1-.743-.648L7 12a.75.75 0 0 1 .649-.743l.101-.007h6.69l-2.72-2.72a.75.75 0 0 1-.072-.976l.073-.084l-.073.084Z"/>`),
		g.Group(children),
	)
}