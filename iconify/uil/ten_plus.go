package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TenPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2Zm-7 4v6a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3Zm5 0v6a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1Zm5.6 0a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9ZM7 7v10a1 1 0 0 0 2 0V7a1 1 0 0 0-2 0Z"/>`),
		g.Group(children),
	)
}