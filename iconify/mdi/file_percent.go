package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePercent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 2l6 6v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8M7.37 20L15 12.35L13.65 11L6 18.65L7.37 20M13 9h5.5L13 3.5V9m-5.5 2A1.5 1.5 0 0 0 6 12.5A1.5 1.5 0 0 0 7.5 14A1.5 1.5 0 0 0 9 12.5A1.5 1.5 0 0 0 7.5 11m6 6a1.5 1.5 0 0 0-1.5 1.5a1.5 1.5 0 0 0 1.5 1.5a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}