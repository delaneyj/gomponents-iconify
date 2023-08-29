package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 4.5v4.967l.991-1.717A3.5 3.5 0 0 1 5.022 6h7.928a2.5 2.5 0 0 0-2.45-2H7.207l-1.56-1.56A1.5 1.5 0 0 0 4.585 2H3.5A2.5 2.5 0 0 0 1 4.5ZM5.022 7h7.973c1.54 0 2.502 1.667 1.732 3l-1.585 2.745a2.5 2.5 0 0 1-2.165 1.25H3.004c-1.54 0-2.502-1.666-1.732-3L2.857 8.25A2.5 2.5 0 0 1 5.022 7Z"/>`),
		g.Group(children),
	)
}