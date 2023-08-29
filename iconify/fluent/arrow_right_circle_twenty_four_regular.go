package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightCircleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12 2c5.524 0 10 4.478 10 10c0 5.524-4.476 10-10 10c-5.522 0-10-4.476-10-10C2 6.479 6.479 2 12 2zm0 1.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17zm-.352 4.054l.072-.084a.75.75 0 0 1 .977-.073l.084.073l4 4a.75.75 0 0 1 .073.977l-.072.084l-4.002 4a.75.75 0 0 1-1.133-.977l.073-.084l2.722-2.72H7.75a.75.75 0 0 1-.743-.648L7 12a.75.75 0 0 1 .648-.743l.102-.007h6.69l-2.72-2.72a.75.75 0 0 1-.072-.976l.072-.084l-.072.084z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}