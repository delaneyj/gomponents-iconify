package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 3v3.25A2.75 2.75 0 0 0 15.75 9H25v12.25A3.75 3.75 0 0 1 21.25 25H6.75A3.75 3.75 0 0 1 3 21.25V6.75A3.75 3.75 0 0 1 6.75 3H13Zm1.5 0v3.25c0 .69.56 1.25 1.25 1.25H25v-.75A3.75 3.75 0 0 0 21.25 3H14.5Z"/>`),
		g.Group(children),
	)
}