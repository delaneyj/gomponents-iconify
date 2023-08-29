package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16.25H5a.74.74 0 0 1-.69-.46a.75.75 0 0 1 .16-.79l7-7a.75.75 0 0 1 1.06 0l7 7a.75.75 0 0 1 .16.82a.74.74 0 0 1-.69.43Zm-12.19-1.5h10.38L12 9.56Z"/>`),
		g.Group(children),
	)
}