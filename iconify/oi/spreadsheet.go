package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.75 0C.34 0 0 .34 0 .75v5.5c0 .41.34.75.75.75h6.5c.41 0 .75-.34.75-.75V.75C8 .34 7.66 0 7.25 0H.75zM1 1h1v1H1V1zm2 0h4v1H3V1zM1 3h1v1H1V3zm2 0h4v1H3V3zM1 5h1v1H1V5zm2 0h4v1H3V5z"/>`),
		g.Group(children),
	)
}