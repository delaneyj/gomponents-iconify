package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FurnitureFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9 10.142v-1.64A.501.501 0 0 0 8.499 8H7.5A.501.501 0 0 0 7 8.501v1.64a3.993 3.993 0 0 0-2.957 3.273a.507.507 0 0 0 .496.586h6.922c.31 0 .541-.28.496-.586A3.993 3.993 0 0 0 9 10.142z" fill="currentColor"/><path d="M13.64 6.279l-2.502-5.004A.498.498 0 0 0 10.692 1H5.308a.498.498 0 0 0-.446.276L2.361 6.279A.498.498 0 0 0 2.806 7H11v1.5a.5.5 0 0 0 1 0V7h1.194c.37 0 .611-.39.445-.721z" fill="currentColor"/>`),
		g.Group(children),
	)
}