package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartVennDiagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 6a9.92 9.92 0 0 0-4 .84A9.92 9.92 0 0 0 12 6a10 10 0 0 0 0 20a9.92 9.92 0 0 0 4-.84a9.92 9.92 0 0 0 4 .84a10 10 0 0 0 0-20Zm-8 18a8 8 0 0 1 0-16a7.91 7.91 0 0 1 1.76.2a10 10 0 0 0 0 15.6A7.91 7.91 0 0 1 12 24Zm8-8a8 8 0 0 1-4 6.92a8 8 0 0 1 0-13.84A8 8 0 0 1 20 16Zm0 8a7.91 7.91 0 0 1-1.76-.2a10 10 0 0 0 0-15.6A7.91 7.91 0 0 1 20 8a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}