package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSZoomIn0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M21 38c9.389 0 17-7.611 17-17S30.389 4 21 4S4 11.611 4 21s7.611 17 17 17Z"/><path stroke="#000" stroke-linecap="round" d="M21 15v12m-5.984-5.984L27 21"/><path stroke="#fff" stroke-linecap="round" d="m33.222 33.222l8.485 8.485"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSZoomIn0)"/>`),
		g.Group(children),
	)
}