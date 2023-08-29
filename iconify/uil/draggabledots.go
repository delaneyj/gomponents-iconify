package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Draggabledots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 10a2 2 0 1 0 2 2a2 2 0 0 0-2-2Zm0 7a2 2 0 1 0 2 2a2 2 0 0 0-2-2Zm7-10a2 2 0 1 0-2-2a2 2 0 0 0 2 2Zm-7-4a2 2 0 1 0 2 2a2 2 0 0 0-2-2Zm7 14a2 2 0 1 0 2 2a2 2 0 0 0-2-2Zm0-7a2 2 0 1 0 2 2a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}