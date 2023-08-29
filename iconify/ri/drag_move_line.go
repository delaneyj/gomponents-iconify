package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragMoveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2l4.243 4.243l-1.415 1.414L12 4.828L9.172 7.657L7.757 6.243L12 2ZM2 12l4.243-4.243l1.414 1.415L4.828 12l2.829 2.828l-1.414 1.415L2 12Zm20 0l-4.243 4.243l-1.414-1.415L19.172 12l-2.829-2.828l1.414-1.415L22 12Zm-10 2a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 8l-4.243-4.243l1.415-1.414L12 19.172l2.828-2.829l1.415 1.414L12 22Z"/>`),
		g.Group(children),
	)
}