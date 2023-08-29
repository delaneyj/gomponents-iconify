package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSortTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6.288 4.293l-3.995 4l-.084.095a1 1 0 0 0 .084 1.32l.095.083a1 1 0 0 0 1.32-.084L6 7.41V19l.007.117a1 1 0 0 0 .993.884l.117-.007A1 1 0 0 0 8 19V7.417l2.293 2.29l.095.084a1 1 0 0 0 1.319-1.499l-4.006-4l-.094-.083a1 1 0 0 0-1.32.084ZM17 4.003l-.117.007a1 1 0 0 0-.883.993v11.58l-2.293-2.29l-.095-.084a1 1 0 0 0-1.319 1.498l4.004 4l.094.084a1 1 0 0 0 1.32-.084l3.996-4l.084-.095a1 1 0 0 0-.084-1.32l-.095-.083a1 1 0 0 0-1.32.084L18 16.587V5.003l-.007-.116A1 1 0 0 0 17 4.003Z"/>`),
		g.Group(children),
	)
}