package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightCircleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12 2c5.524 0 10 4.478 10 10c0 5.524-4.476 10-10 10c-5.522 0-10-4.476-10-10C2 6.479 6.479 2 12 2zm.781 5.47l-.084-.073a.75.75 0 0 0-.883-.007l-.094.08l-.072.084a.75.75 0 0 0-.007.882l.08.094l2.719 2.72H7.75l-.102.007a.75.75 0 0 0-.641.641L7 12l.007.102a.75.75 0 0 0 .641.641l.102.007h6.69l-2.72 2.72l-.073.084a.75.75 0 0 0 1.05 1.05l.083-.073l4.002-4l.072-.084a.75.75 0 0 0 .008-.883l-.08-.094l-4.001-4l-.084-.073l.084.073z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}