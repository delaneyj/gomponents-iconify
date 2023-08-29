package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 22.288h2V7.802l3.243 3.243l1.414-1.414L12 3.974L6.343 9.63l1.414 1.414L11 7.802v14.486ZM18 3H6V1h12v2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}