package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutgoingMailOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 20l-1.4-1.4l1.575-1.6H14v-2h4.175L16.6 13.4L18 12l4 4l-4 4ZM4 17q-.825 0-1.413-.588T2 15V5q0-.825.588-1.413T4 3h13q.825 0 1.413.588T19 5v5.1q-.25-.05-.5-.075T18 10q-.25 0-.5.013t-.5.062V6.4L10.4 11L4 6.425V15h8.075q-.05.25-.062.5T12 16q0 .25.025.5t.075.5H4ZM5.45 5l4.95 3.55L15.5 5H5.45ZM4 15V5v10Z"/>`),
		g.Group(children),
	)
}