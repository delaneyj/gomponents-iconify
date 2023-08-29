package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterDismissTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.726 10.25a5.46 5.46 0 0 0 2.274.727V11a.75.75 0 0 1-.75.75h-8.5a.75.75 0 0 1 0-1.5h6.976ZM9.022 6c.048.525.169 1.028.353 1.5H2.75a.75.75 0 0 1 0-1.5h6.272ZM12 15.25a.75.75 0 0 0-.75-.75h-4.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 .75-.75ZM14.5 10a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm1.854-6.354a.5.5 0 0 1 0 .708L15.207 5.5l1.147 1.146a.5.5 0 0 1-.708.708L14.5 6.207l-1.146 1.147a.5.5 0 0 1-.708-.708L13.793 5.5l-1.147-1.146a.5.5 0 0 1 .708-.708L14.5 4.793l1.146-1.147a.5.5 0 0 1 .708 0Z"/>`),
		g.Group(children),
	)
}