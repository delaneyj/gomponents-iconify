package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeZigzag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.312 9L1 10.51l3.774 3.28l1.509 1.312l1.312-1.51l1.54-1.77l2.264 1.968l1.51 1.312l1.311-1.51l1.538-1.769l2.263 1.967l1.51 1.312l1.311-1.51l1.969-2.264l-1.51-1.312l-1.968 2.264L15.559 9l-1.312 1.51h.002l-1.538 1.77L8.937 9l-.883 1.016l-1.968 2.264L2.312 9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}