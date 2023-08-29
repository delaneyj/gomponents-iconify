package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.646 7.146a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 1 1-.708.708L6.5 8.707V11.5a.5.5 0 0 1-1 0V8.707l-.646.647a.5.5 0 1 1-.708-.708l1.5-1.5Zm10.208 3.208l-1.5 1.5a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 0 1 .708-.708l.646.647V7.5a.5.5 0 0 1 1 0v2.793l.646-.647a.5.5 0 0 1 .708.708ZM5 3a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H5Zm4.5 1v11H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h4.5Zm1 11V4H15a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2h-4.5Z"/>`),
		g.Group(children),
	)
}