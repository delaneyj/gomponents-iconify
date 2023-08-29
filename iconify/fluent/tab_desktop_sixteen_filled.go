package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V5H8.5A1.5 1.5 0 0 1 7 3.5V2Zm7 2a2 2 0 0 0-2-2H8v1.5a.5.5 0 0 0 .5.5H14Z"/>`),
		g.Group(children),
	)
}