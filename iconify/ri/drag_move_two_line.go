package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragMoveTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11V5.828L9.172 7.657L7.757 6.243L12 2l4.243 4.243l-1.415 1.414L13 5.828V11h5.172l-1.829-1.828l1.414-1.415L22 12l-4.243 4.243l-1.414-1.415L18.172 13H13v5.172l1.828-1.829l1.415 1.414L12 22l-4.243-4.243l1.415-1.414L11 18.172V13H5.828l1.829 1.828l-1.414 1.415L2 12l4.243-4.243l1.414 1.415L5.828 11H11Z"/>`),
		g.Group(children),
	)
}