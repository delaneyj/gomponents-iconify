package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.929 5l1.414 1.414l-5.657 5.657l5.657 5.657l-1.414 1.414l-7.071-7.07L16.929 5ZM8 19V5H6v14h2Z"/>`),
		g.Group(children),
	)
}