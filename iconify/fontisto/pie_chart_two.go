package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 11.981C.004 5.612 4.965.403 11.233.002L11.268 0v11.981c0 .139.038.269.105.379l-.002-.003l5.989 10.378A11.831 11.831 0 0 1 12 23.999c-6.627 0-12-5.373-12-12v-.018v.001zm13.318.752H24a12.006 12.006 0 0 1-5.291 9.231l-.043.027zm-.548-1.503V0c6.04.388 10.842 5.19 11.229 11.194l.002.035z"/>`),
		g.Group(children),
	)
}