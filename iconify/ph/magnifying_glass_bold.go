package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232.49 215.51L185 168a92.12 92.12 0 1 0-17 17l47.53 47.54a12 12 0 0 0 17-17ZM44 112a68 68 0 1 1 68 68a68.07 68.07 0 0 1-68-68Z"/>`),
		g.Group(children),
	)
}