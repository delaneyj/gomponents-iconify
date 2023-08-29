package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuCake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-7 2a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H5Zm-1 5a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1Zm1 3a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H5Z"/>`),
		g.Group(children),
	)
}