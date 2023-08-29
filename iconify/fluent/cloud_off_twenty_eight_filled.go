package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l6.084 6.085A6.46 6.46 0 0 0 7.52 12h-.27a4.75 4.75 0 1 0 0 9.5h13.189l4.28 4.28a.75.75 0 0 0 1.061-1.06L3.28 2.22Zm7.041 4.92l13.353 13.353A4.75 4.75 0 0 0 20.75 12h-.269a6.5 6.5 0 0 0-10.16-4.86Z"/>`),
		g.Group(children),
	)
}