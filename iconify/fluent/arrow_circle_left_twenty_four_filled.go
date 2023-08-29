package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Zm.28 14.53a.75.75 0 0 1-.976.073l-.085-.072l-4-4.001a.75.75 0 0 1-.073-.977l.073-.084l4.001-4a.75.75 0 0 1 1.133.977l-.072.084l-2.722 2.72h6.691a.75.75 0 0 1 .744.649L17 12a.75.75 0 0 1-.648.743l-.102.007H9.56l2.72 2.72a.75.75 0 0 1 .073.977l-.073.084Z"/>`),
		g.Group(children),
	)
}