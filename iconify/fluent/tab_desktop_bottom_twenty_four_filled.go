package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopBottomTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.25A3.25 3.25 0 0 1 6.25 3h11.5A3.25 3.25 0 0 1 21 6.25v11.5A3.25 3.25 0 0 1 17.75 21H14.5v-3.25a2.25 2.25 0 0 0-2.25-2.25H3V6.25ZM3 17v.75A3.25 3.25 0 0 0 6.25 21H13v-3.25a.75.75 0 0 0-.75-.75H3Z"/>`),
		g.Group(children),
	)
}