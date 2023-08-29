package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8h-7.75A2.25 2.25 0 0 1 11 5.75V3Zm10 3.5V6a3 3 0 0 0-3-3h-5.5v2.75c0 .414.336.75.75.75H21Z"/>`),
		g.Group(children),
	)
}