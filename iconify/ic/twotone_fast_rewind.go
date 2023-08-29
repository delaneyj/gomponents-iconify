package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneFastRewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14.14V9.86L5.97 12zm9 0V9.86L14.97 12z" opacity=".3"/><path fill="currentColor" d="m11 6l-8.5 6l8.5 6V6zm-2 8.14L5.97 12L9 9.86v4.28zM20 6l-8.5 6l8.5 6V6zm-2 8.14L14.97 12L18 9.86v4.28z"/>`),
		g.Group(children),
	)
}