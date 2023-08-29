package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.49 9.5H8.75a.75.75 0 0 1 0-1.5h6.5a.75.75 0 0 1 .75.75v6.5a.75.75 0 0 1-1.5 0v-4.639l-5.222 5.172a.75.75 0 0 1-1.056-1.066L13.491 9.5ZM2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Zm10 8.5a8.5 8.5 0 1 1 0-17a8.5 8.5 0 0 1 0 17Z"/>`),
		g.Group(children),
	)
}