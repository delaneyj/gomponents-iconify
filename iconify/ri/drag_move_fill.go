package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragMoveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-4-4h8l-4 4Zm0-20l4 4H8l4-4Zm0 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM2 12l4-4v8l-4-4Zm20 0l-4 4V8l4 4Z"/>`),
		g.Group(children),
	)
}