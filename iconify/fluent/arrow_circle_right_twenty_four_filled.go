package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.001 2c5.523 0 10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10Zm.78 5.469l-.084-.073a.75.75 0 0 0-.882-.007l-.094.08l-.073.084a.75.75 0 0 0-.007.883l.08.094l2.72 2.72H7.75l-.103.006a.75.75 0 0 0-.64.642L7 11.999l.007.102a.75.75 0 0 0 .642.641l.101.007h6.69l-2.72 2.72l-.073.085a.75.75 0 0 0 1.05 1.05l.084-.073l4.001-4l.073-.085a.75.75 0 0 0 .007-.882l-.08-.094l-4-4.001l-.085-.073l.084.073Z"/>`),
		g.Group(children),
	)
}