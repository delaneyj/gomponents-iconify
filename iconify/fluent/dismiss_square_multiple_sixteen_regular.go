package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DismissSquareMultipleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.646 4.646a.5.5 0 0 1 .708 0L7 6.293l1.646-1.647a.5.5 0 1 1 .708.708L7.707 7l1.647 1.646a.5.5 0 1 1-.708.708L7 7.707L5.354 9.354a.5.5 0 1 1-.708-.708L6.293 7L4.646 5.354a.5.5 0 0 1 0-.708ZM4 2a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4ZM3 4a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4Zm10-.732A2 2 0 0 1 14 5v5.5a3.5 3.5 0 0 1-3.5 3.5H5a2 2 0 0 1-1.732-1H10.5a2.5 2.5 0 0 0 2.5-2.5V3.268Z"/>`),
		g.Group(children),
	)
}