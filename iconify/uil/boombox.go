package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boombox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm1 17a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm-3.62-8.2A3 3 0 0 0 15 9a3 3 0 0 0-6 0a3 3 0 0 0 .62 1.8a4 4 0 1 0 4.76 0ZM12 8a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm0 8a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}