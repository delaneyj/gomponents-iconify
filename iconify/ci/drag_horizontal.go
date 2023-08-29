package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 14a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm12-6a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM6 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}