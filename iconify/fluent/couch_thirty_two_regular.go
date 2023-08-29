package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouchThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 6a4 4 0 0 0-4 4v1.256A4.502 4.502 0 0 0 2 15.5V20a4.002 4.002 0 0 0 3 3.874V26a1 1 0 1 0 2 0v-2h18v2a1 1 0 1 0 2 0v-2.126c1.725-.444 3-2.01 3-3.874v-4.5a4.502 4.502 0 0 0-3-4.244V10a4 4 0 0 0-4-4H9Zm17 16H6a2 2 0 0 1-2-2v-4.5a2.5 2.5 0 0 1 5 0v.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-.5a2.5 2.5 0 0 1 5 0V20a2 2 0 0 1-2 2Zm-1-10.973A4.502 4.502 0 0 0 21.027 15H10.973A4.502 4.502 0 0 0 7 11.027V10a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v1.027Z"/>`),
		g.Group(children),
	)
}