package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 10.75H6a.74.74 0 0 1-.69-.46a.75.75 0 0 1 .16-.82l6-6a.75.75 0 0 1 1.06 0l6 6a.75.75 0 0 1 .16.82a.74.74 0 0 1-.69.46ZM7.81 9.25h8.38L12 5.06ZM12 20.75a.74.74 0 0 1-.53-.22l-6-6a.75.75 0 0 1-.16-.82a.74.74 0 0 1 .69-.46h12a.74.74 0 0 1 .69.46a.75.75 0 0 1-.16.82l-6 6a.74.74 0 0 1-.53.22Zm-4.19-6L12 18.94l4.19-4.19Z"/>`),
		g.Group(children),
	)
}