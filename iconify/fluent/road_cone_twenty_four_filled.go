package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadConeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.194 2a1.25 1.25 0 0 0-1.215.955C9.288 5.797 8.444 8.905 7.594 12h5.156a.75.75 0 0 1 0 1.5H7.181l-.55 2h7.119a.75.75 0 0 1 0 1.5H6.222c-.325 1.198-.639 2.37-.932 3.5H2.75a.75.75 0 0 0 0 1.5h18.5a.75.75 0 0 0 0-1.5h-2.54L14.023 2.928A1.25 1.25 0 0 0 12.816 2h-1.622Z"/>`),
		g.Group(children),
	)
}