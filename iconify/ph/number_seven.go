package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m175.66 50.3l-48 160a8 8 0 0 1-15.32-4.6L157.25 56H88a8 8 0 0 1 0-16h80a8 8 0 0 1 7.66 10.3Z"/>`),
		g.Group(children),
	)
}