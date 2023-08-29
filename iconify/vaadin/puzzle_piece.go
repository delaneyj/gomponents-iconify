package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PuzzlePiece(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.9.9c-1.1-1-2.5-1.3-3.1-.4c-.7 1.1.5 1.7-.3 2.5c-.5.6-2-.8-2-.8l-.8-.8l-1.4 1.4c-.6.7-2.1 1.5-2.6 1.1c-.7-.6.1-1.8-.5-2.6c-.7-1-2.1-.8-3 .3c-1 1.1-1.4 2.4-.5 3c1.1.7 1.9-.3 2.7.5c.4.4-.2 1.7-.5 2.1L.6 9.5L7.1 16l1.7-1.7c.7-.7 1.5-2 1.1-2.4c-.6-.7-1.7.1-2.5-.4c-1-.7-.8-2 .3-3s2.5-1.3 3.1-.4c.7 1.1-.4 1.8.4 2.6c.4.4 1.6-.2 2-.6L15.3 8l-1.1-1.1c-.6-.6-1.9-2-1.4-2.5c.6-.7 1.7.2 2.5-.4c.9-.8.6-2.1-.4-3.1z"/>`),
		g.Group(children),
	)
}