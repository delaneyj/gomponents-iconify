package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineFastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 9.86L18.03 12L15 14.14V9.86m-9 0L9.03 12L6 14.14V9.86M13 6v12l8.5-6L13 6zM4 6v12l8.5-6L4 6z"/>`),
		g.Group(children),
	)
}