package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadingListSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 4.386a.614.614 0 0 1 1.126-.34a.75.75 0 0 0 1.248-.833A2.113 2.113 0 0 0 2 4.386c0 1.073.799 1.959 1.834 2.096A.753.753 0 0 0 4 6.5h7.31a.75.75 0 0 0 0-1.5H4.114a.614.614 0 0 1-.614-.614ZM7.75 2.5a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5h-5.5Zm-3 5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5ZM2 10.75a.75.75 0 0 1 .75-.75h8.5a.75.75 0 1 1 0 1.5h-8.5a.75.75 0 0 1-.75-.75Zm2.745 1.751a.75.75 0 1 0 0 1.5h8.5a.75.75 0 1 0 0-1.5h-8.5Z"/>`),
		g.Group(children),
	)
}