package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterdropLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 13.193C3 18.057 6.855 22 11.611 22h.777C17.145 22 21 18.057 21 13.193v-.265c0-4.611-2.729-8.765-6.903-10.507a5.434 5.434 0 0 0-4.194 0C5.73 4.163 3 8.317 3 12.928v.265Z" opacity=".5"/><path stroke-linecap="round" d="M7.615 10.724c.634-2.006 2.016-3.636 3.77-4.448"/></g>`),
		g.Group(children),
	)
}