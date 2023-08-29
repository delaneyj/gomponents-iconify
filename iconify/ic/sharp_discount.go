package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpDiscount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.79 21L3 11.21v2.83l9.79 9.79l9.04-9.04l-1.42-1.41z"/><path fill="currentColor" d="m3 9.04l9.79 9.79l9.04-9.04L12.04 0H3v9.04zM7.25 3a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5z"/>`),
		g.Group(children),
	)
}