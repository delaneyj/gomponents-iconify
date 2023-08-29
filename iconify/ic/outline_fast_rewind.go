package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineFastRewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9.86v4.28L14.97 12L18 9.86m-9 0v4.28L5.97 12L9 9.86M20 6l-8.5 6l8.5 6V6zm-9 0l-8.5 6l8.5 6V6z"/>`),
		g.Group(children),
	)
}