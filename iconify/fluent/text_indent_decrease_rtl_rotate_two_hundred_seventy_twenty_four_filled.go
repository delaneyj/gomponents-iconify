package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentDecreaseRtlRotateTwoHundredSeventyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 18.5a1 1 0 0 0 1-1V9l-.007-.117A1 1 0 0 0 6 9v8.5l.007.117A1 1 0 0 0 7 18.5Zm6.001 2L13 9l-.007-.117A1 1 0 0 0 11 9l.001 11.5l.007.117A1 1 0 0 0 13 20.5Zm3.999-2a1 1 0 0 0 1-1V9l-.007-.117A1 1 0 0 0 16 9v8.5l.007.117A1 1 0 0 0 17 18.5ZM14.79 4.387a1 1 0 0 1-1.497 1.32L12 4.414l-1.293 1.293l-.094.083a1 1 0 0 1-1.32-1.497l2-2l.094-.083a1 1 0 0 1 1.32.083l2 2l.083.094Z"/>`),
		g.Group(children),
	)
}