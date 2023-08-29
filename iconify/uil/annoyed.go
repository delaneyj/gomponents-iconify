package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Annoyed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm6 4a5 5 0 0 0-4.37 2.57a1 1 0 0 0 .37 1.36a1 1 0 0 0 .49.13a1 1 0 0 0 .87-.51A3 3 0 0 1 15 15a1 1 0 0 0 0-2Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}