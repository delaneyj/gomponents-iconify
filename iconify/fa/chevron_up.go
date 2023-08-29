package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1619 1075l-166 165q-19 19-45 19t-45-19L832 709l-531 531q-19 19-45 19t-45-19L45 1075q-19-19-19-45.5T45 984l742-741q19-19 45-19t45 19l742 741q19 19 19 45.5t-19 45.5z"/>`),
		g.Group(children),
	)
}