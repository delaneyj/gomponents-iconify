package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.49 14.5H8.75a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 .75-.75v-6.5a.75.75 0 0 0-1.5 0v4.639L9.278 8.217a.75.75 0 1 0-1.056 1.066l5.269 5.217ZM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm10-8.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17Z"/>`),
		g.Group(children),
	)
}