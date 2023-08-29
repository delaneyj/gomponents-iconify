package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2H8a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM6 6a2 2 0 0 1 2-2h10v10H8a3.91 3.91 0 0 0-2 .56Zm2 14a2 2 0 0 1 0-4h10v4Zm2-12h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}