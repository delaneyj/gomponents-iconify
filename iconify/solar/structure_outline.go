package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StructureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 8.675A3.751 3.751 0 1 1 8.675 4.25h6.65a3.751 3.751 0 1 1 4.425 4.425v6.65a3.751 3.751 0 1 1-4.425 4.425h-6.65a3.751 3.751 0 1 1-4.425-4.425v-6.65ZM2.75 5a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Zm3 10.325v-6.65A3.754 3.754 0 0 0 8.675 5.75h6.65a3.755 3.755 0 0 0 2.925 2.925v6.65a3.755 3.755 0 0 0-2.925 2.925h-6.65a3.755 3.755 0 0 0-2.925-2.925ZM5 16.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM21.25 5a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0Zm-4.5 14a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}