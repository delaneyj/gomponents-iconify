package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 7A7.74 7.74 0 1 0 7 17A7.74 7.74 0 1 0 17 7Zm-1.6 1.6a6.12 6.12 0 0 1 .11 1.14a5.92 5.92 0 0 1-.16 1.34l-2.44-2.42a5.92 5.92 0 0 1 1.34-.16a6.12 6.12 0 0 1 1.14.11Zm-.94 4.4A6 6 0 0 1 13 14.46L9.54 11A6 6 0 0 1 11 9.54Zm-8 1.46a5.75 5.75 0 1 1 8-8h-.25a7.76 7.76 0 0 0-7.71 7.79c0 .08.01.17.01.25Zm2.1.89a6.12 6.12 0 0 1-.11-1.14a5.92 5.92 0 0 1 .16-1.34l2.43 2.43a5.92 5.92 0 0 1-1.34.16a6.12 6.12 0 0 1-1.09-.07ZM14.25 20a5.77 5.77 0 0 1-4.75-2.51h.25a7.76 7.76 0 0 0 7.75-7.74V9.5A5.75 5.75 0 0 1 14.25 20Z"/>`),
		g.Group(children),
	)
}