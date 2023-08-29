package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 2v2.5A1.5 1.5 0 0 0 6.5 6h2A1.5 1.5 0 0 0 10 4.5V2h.379a2 2 0 0 1 1.414.586l1.621 1.621A2 2 0 0 1 14 5.621V12a2 2 0 0 1-2 2V9.5A1.5 1.5 0 0 0 10.5 8h-5C4.673 8 4 8.669 4 9.498V14a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h1Zm1 0v2.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V2H6ZM5 14h6V9.5a.5.5 0 0 0-.5-.5h-5c-.277 0-.5.223-.5.498V14Z"/>`),
		g.Group(children),
	)
}