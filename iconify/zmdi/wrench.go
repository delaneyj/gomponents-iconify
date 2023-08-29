package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M464 389q6 5 6 14.5t-8 15.5l-49 49q-7 7-15.5 7t-14.5-7L189 274q-37 15-77.5 6.5T41 242q-31-32-39-75.5T14 84l94 92l64-64l-92-92q38-18 82-10.5T238 48q30 30 38.5 70.5T270 195z"/>`),
		g.Group(children),
	)
}