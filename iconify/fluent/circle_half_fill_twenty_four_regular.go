package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfFillTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.028 11.25A9.999 9.999 0 0 1 12 2c5.27 0 9.589 4.077 9.972 9.25H22V12c0 5.523-4.477 10-10 10S2 17.523 2 12v-.75h.028ZM12 3.5a8.5 8.5 0 0 0-8.467 7.75h16.934A8.5 8.5 0 0 0 12 3.5Z"/>`),
		g.Group(children),
	)
}