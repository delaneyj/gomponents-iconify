package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 16.929l1.414 1.414l5.657-5.657l5.657 5.657l1.414-1.414l-7.07-7.071L5 16.929ZM19 8H5V6h14v2Z"/>`),
		g.Group(children),
	)
}