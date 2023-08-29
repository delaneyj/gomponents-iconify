package svg_spinners

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDotsRotate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="12" r="3" fill="currentColor"/><g><circle cx="4" cy="12" r="3" fill="currentColor"/><circle cx="20" cy="12" r="3" fill="currentColor"/><animateTransform attributeName="transform" calcMode="spline" dur="1s" keySplines=".36,.6,.31,1;.36,.6,.31,1" repeatCount="indefinite" type="rotate" values="0 12 12;180 12 12;360 12 12"/></g>`),
		g.Group(children),
	)
}