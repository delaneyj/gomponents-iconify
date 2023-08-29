package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneFastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 9.86v4.28L18.03 12zm-9 0v4.28L9.03 12z" opacity=".3"/><path fill="currentColor" d="m4 18l8.5-6L4 6v12zm2-8.14L9.03 12L6 14.14V9.86zM21.5 12L13 6v12l8.5-6zM15 9.86L18.03 12L15 14.14V9.86z"/>`),
		g.Group(children),
	)
}