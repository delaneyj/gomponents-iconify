package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudMeatball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 15.92h-.77l.39-.67a1 1 0 0 0-1.74-1l-.38.67l-.38-.67a1 1 0 0 0-1.74 1l.39.67H9.5a1 1 0 0 0 0 2h.77l-.39.66a1 1 0 0 0 1.74 1l.38-.66l.38.66a1 1 0 0 0 1.74-1l-.39-.66h.77a1 1 0 0 0 0-2Zm3.92-7.79A7 7 0 0 0 5.06 10A4 4 0 0 0 2 13.92a4 4 0 0 0 3.34 3.93h.16a1 1 0 0 0 .16-2a2 2 0 0 1-1.66-2a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.66a3 3 0 0 1 .62 5.72a1 1 0 1 0 .74 1.85a5 5 0 0 0-.45-9.41Z"/>`),
		g.Group(children),
	)
}