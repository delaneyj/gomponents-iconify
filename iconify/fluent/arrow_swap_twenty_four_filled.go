package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSwapTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m15.207 2.29l4 3.996a1 1 0 0 1 .084 1.32l-.083.094l-4 4.006a1 1 0 0 1-1.498-1.32l.083-.094L16.083 8H5.5a1 1 0 0 1-.993-.883L4.5 6.999a1 1 0 0 1 .883-.993l.117-.007h10.59l-2.296-2.293a1 1 0 0 1-.084-1.32l.083-.095a1 1 0 0 1 1.32-.084l.094.084l4 3.995l-4-3.995Zm4.284 14.592l.006.117a1 1 0 0 1-.883.993l-.117.007H7.914l2.293 2.293a1 1 0 0 1 .084 1.32l-.083.094a1 1 0 0 1-1.32.084l-.094-.084l-4-3.996a1 1 0 0 1-.084-1.32l.083-.094l4-4.004a1 1 0 0 1 1.498 1.32l-.083.094l-2.29 2.293h10.58a1 1 0 0 1 .993.883l.006.117l-.006-.117Z"/>`),
		g.Group(children),
	)
}