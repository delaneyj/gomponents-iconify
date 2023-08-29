package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRadar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 .764L23.186 8.89l-4.273 13.15H5.087L.814 8.89L12 .764ZM3.548 10.83l2.61 8.035l4.224-5.814l-6.834-2.22Zm4.228 9.21h8.448L12 14.226L7.776 20.04Zm5.842-6.99l4.224 5.815l2.61-8.035l-6.834 2.22Zm6.216-4.122L13 3.963v7.186l6.834-2.22ZM11 3.963L4.166 8.928L11 11.148V3.964Z"/>`),
		g.Group(children),
	)
}