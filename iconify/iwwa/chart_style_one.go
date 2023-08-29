package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStyleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M30.773 31.003H7.29l3.551-8.522l2.625 2.017l4.451-7.158l5.009 3.082l7.847-16.568v27.149zM9.038 29.838h20.57V9.036l-6.186 13.061l-5.128-3.158l-4.517 7.267l-2.443-1.879l-2.296 5.511zm-.68 3.509H30.19v1.165H8.358z"/>`),
		g.Group(children),
	)
}