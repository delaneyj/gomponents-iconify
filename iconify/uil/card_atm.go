package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardAtm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4.5H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3Zm1 12a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-4-6a3 3 0 0 0-1.51.42a3 3 0 1 0 0 5.16A3 3 0 1 0 16 10.5Zm-2.83 4a1 1 0 0 1-.17 0a1 1 0 0 1 0-2a1 1 0 0 1 .17 0a2.8 2.8 0 0 0 0 1.92Zm2.83 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}