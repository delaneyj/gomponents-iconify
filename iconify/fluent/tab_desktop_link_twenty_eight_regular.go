package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopLinkTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75v14.5A3.75 3.75 0 0 0 6.75 25h6.583a4.727 4.727 0 0 1-.326-1.5H6.75a2.25 2.25 0 0 1-2.25-2.25V6.75A2.25 2.25 0 0 1 6.75 4.5H13v1.75A2.75 2.75 0 0 0 15.75 9h7.75v9.506a4.737 4.737 0 0 1 1.5.327V6.75A3.75 3.75 0 0 0 21.25 3H6.75ZM23.5 7.5h-7.75c-.69 0-1.25-.56-1.25-1.25V4.5h6.75a2.25 2.25 0 0 1 2.25 2.25v.75ZM17.75 21a2.25 2.25 0 0 0 0 4.5h.5a.75.75 0 0 1 0 1.5h-.5a3.75 3.75 0 1 1 0-7.5h.5a.75.75 0 0 1 0 1.5h-.5ZM17 23.25a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5a.75.75 0 0 1-.75-.75Zm6.25 2.25a2.25 2.25 0 0 0 0-4.5h-.5a.75.75 0 0 1 0-1.5h.5a3.75 3.75 0 1 1 0 7.5h-.5a.75.75 0 0 1 0-1.5h.5Z"/>`),
		g.Group(children),
	)
}