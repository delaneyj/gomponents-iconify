package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopArrowClockwiseTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.75 3A2.75 2.75 0 0 0 3 5.75v5.351a7.029 7.029 0 0 1 1.5-1.165V5.75c0-.69.56-1.25 1.25-1.25H11v1.25A2.25 2.25 0 0 0 13.25 8h6.25v10.25c0 .69-.56 1.25-1.25 1.25h-4.187A7.028 7.028 0 0 1 12.9 21h5.351A2.75 2.75 0 0 0 21 18.25V5.75A2.75 2.75 0 0 0 18.25 3H5.75ZM19.5 6.5h-6.25a.75.75 0 0 1-.75-.75V4.5h5.75c.69 0 1.25.56 1.25 1.25v.75Zm-7.25 3.25a.75.75 0 0 1 .75.75v2.75a.75.75 0 0 1-.75.75H9.5a.75.75 0 0 1 0-1.5h1.33a4.478 4.478 0 0 0-2.83-1a4.5 4.5 0 1 0 4.5 4.5a.75.75 0 0 1 1.5 0a6 6 0 1 1-2.5-4.874V10.5a.75.75 0 0 1 .75-.75Z"/>`),
		g.Group(children),
	)
}