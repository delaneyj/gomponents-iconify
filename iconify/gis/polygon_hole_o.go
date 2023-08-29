package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolygonHoleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M33.39 20.396a3.5 3.496 0 0 0-3.11 1.971L9.573 65.117a3.5 3.496 0 0 0 2.631 4.979l75.002 11.261a3.5 3.496 0 0 0 3.844-4.544L76.787 33.374a3.5 3.496 0 0 0-2.432-2.291l-40.03-10.572a3.5 3.496 0 0 0-.934-.116zm1.932 7.612l35.407 9.35L82.63 73.601l-64.688-9.713l17.38-35.881zm8.033 11.324a2.5 2.5 0 0 0-2.332 1.512l-5.95 13.826a2.5 2.5 0 0 0 1.888 3.455l24.328 4.025a2.5 2.5 0 0 0 2.799-1.738l2.451-8.05a2.5 2.5 0 0 0-1.328-2.99L44.383 39.57a2.5 2.5 0 0 0-1.028-.238zm1.239 5.863l16.531 7.78l-1.182 3.884l-19.015-3.146l3.666-8.518z" color="currentColor"/>`),
		g.Group(children),
	)
}