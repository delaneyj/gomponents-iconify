package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm4 0h7a3 3 0 0 0 0-6a3.06 3.06 0 0 0-1.5.4a1 1 0 0 0-.37 1.37a1 1 0 0 0 1.37.36a1.09 1.09 0 0 1 .5-.13a1 1 0 0 1 0 2h-7a1 1 0 0 0 0 2Zm-4 4h7a1 1 0 0 0 0-2h-7a1 1 0 0 0 0 2Zm17-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-2 2h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 0 0 0-6Zm-6 4h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 0 2a1.09 1.09 0 0 1-.5-.13a1 1 0 1 0-1 1.73a3.06 3.06 0 0 0 1.5.4a3 3 0 0 0 0-6Zm-8 0h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}