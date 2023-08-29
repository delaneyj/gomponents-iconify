package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentDecreaseRtlNinetyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 5.5a1 1 0 0 0-1 1V15l.007.117A1 1 0 0 0 18 15V6.5l-.007-.117A1 1 0 0 0 17 5.5Zm-6.001-2L11 15l.007.117A1 1 0 0 0 13 15l-.001-11.5l-.007-.117A1 1 0 0 0 11 3.5ZM7 5.5a1 1 0 0 0-1 1V15l.007.117A1 1 0 0 0 8 15V6.5l-.007-.117A1 1 0 0 0 7 5.5Zm2.21 14.113a1 1 0 0 1 1.497-1.32L12 19.586l1.293-1.293l.094-.083a1 1 0 0 1 1.32 1.497l-2 2l-.094.083a1 1 0 0 1-1.32-.083l-2-2l-.083-.094Z"/>`),
		g.Group(children),
	)
}