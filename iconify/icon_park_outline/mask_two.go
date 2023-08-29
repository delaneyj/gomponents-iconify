package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaskTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="24" r="20"/><path stroke-linecap="round" stroke-linejoin="round" d="M31 5.26C23.407 8.096 18 15.416 18 24c0 8.583 5.407 15.903 13 18.74M37 9L18 22m23-8L19 29m24-9L22 35m21-7L26 40"/></g>`),
		g.Group(children),
	)
}