package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenVerticalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 4.5A2.5 2.5 0 0 0 14.5 2H7.529l2.716 1.568a3.5 3.5 0 0 1 1.75 3.031V16h.505a2.5 2.5 0 0 0 2.5-2.5V9.707l1.56-1.56A1.5 1.5 0 0 0 17 7.085V4.5Zm-7.255-.066a2.5 2.5 0 0 1 1.25 2.165v10.396c0 1.54-1.667 2.502-3 1.732L4.25 16.565A2.5 2.5 0 0 1 3 14.4V4.004c0-1.54 1.667-2.502 3-1.732l3.745 2.162Z"/>`),
		g.Group(children),
	)
}