package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.998.811l8.384 5.388L19.3 7.882l-.3-.193v3.417l3.375 2.062l-1.043 1.707l-.332-.203V22H3v-7.328l-.332.203l-1.043-1.707L5 11.106V7.689l-.3.193L3.617 6.2L11.998.81ZM7 6.403v3.48l5-3.055l5 3.055v-3.48L11.998 3.19L7 6.403ZM5 13.45V20h6v-4h2v4h6v-6.55l-7-4.278l-7 4.278Z"/>`),
		g.Group(children),
	)
}