package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractSquareMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4Zm.5 4.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 1 1 0-1Zm8.5 4a2.5 2.5 0 0 1-2.5 2.5H3.268A2 2 0 0 0 5 14h5.5a3.5 3.5 0 0 0 3.5-3.5V5a2 2 0 0 0-1-1.732V10.5Z"/>`),
		g.Group(children),
	)
}