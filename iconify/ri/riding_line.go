package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RidingLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 21a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm0-2a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm13 2a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm0-2a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-7.477-8.695L13 12v6h-2v-5l-2.719-2.266A2 2 0 0 1 8 7.671l2.828-2.828a2 2 0 0 1 2.829 0l1.414 1.414a6.969 6.969 0 0 0 3.917 1.975l-.01 2.015a8.962 8.962 0 0 1-5.321-2.575l-2.634 2.633ZM16 5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}