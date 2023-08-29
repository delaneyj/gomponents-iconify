package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.378 3.026A9.004 9.004 0 1 1 5.475 17.13a.675.675 0 0 1 .329-1.019c3.391-1.214 5.208-2.62 6.262-4.633c1.11-2.118 1.396-4.438.62-7.619a.675.675 0 0 1 .692-.834Zm-2.384.627L11 3.75V9a2 2 0 1 1-1.499-1.937l-.001-2.3L6 5.829V10a2 2 0 1 1-1.499-1.937L4.5 5.273a.75.75 0 0 1 .43-.679l.102-.039l5-1.521a.75.75 0 0 1 .962.619Z"/>`),
		g.Group(children),
	)
}