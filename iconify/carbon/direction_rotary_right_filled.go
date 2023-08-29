package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionRotaryRightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="11" cy="11" r="3" fill="currentColor"/><path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-8 16l-1.414-1.414L23.172 12h-7.273A5.014 5.014 0 0 1 12 15.899V26h-2V15.899A5 5 0 1 1 15.899 10h7.273l-4.586-4.586L20 4l7 7Z"/><path fill="none" d="m20 4l-1.414 1.414L23.172 10h-7.273A5 5 0 1 0 10 15.899V26h2V15.899A5.014 5.014 0 0 0 15.899 12h7.273l-4.586 4.586L20 18l7-7Zm-9 10a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z"/>`),
		g.Group(children),
	)
}