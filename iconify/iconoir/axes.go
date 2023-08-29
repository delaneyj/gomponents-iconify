package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 19.452l-9-6.61m0 0V3m0 9.843l-9 6.609m17.438-2.742L21 19.452L18.187 20M9.75 5.194L12 3l2.25 2.194M5.813 20L3 19.452l.563-2.742"/>`),
		g.Group(children),
	)
}