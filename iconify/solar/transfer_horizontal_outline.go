package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferHorizontalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.007 3.447a.75.75 0 0 1 .046 1.06L5.705 9.25H20a.75.75 0 0 1 0 1.5H4a.75.75 0 0 1-.553-1.257l5.5-6a.75.75 0 0 1 1.06-.046Zm10.679 10.251A.75.75 0 0 0 20 13.25H4a.75.75 0 0 0 0 1.5h14.295l-4.348 4.743a.75.75 0 0 0 1.106 1.014l5.5-6a.75.75 0 0 0 .133-.809Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}