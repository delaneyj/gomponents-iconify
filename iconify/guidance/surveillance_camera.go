package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveillanceCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M23.5 9v5m-13-.5v.649a7.5 7.5 0 0 0 1.894 4.982l.106.119v.25H0m20.5-6v-2.7a7.5 7.5 0 0 1 2.432-5.53l.568-.52V4.5H5a4.5 4.5 0 0 0 0 9h15.5Z"/>`),
		g.Group(children),
	)
}