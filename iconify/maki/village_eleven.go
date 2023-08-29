package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VillageEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M2.777 2.3L.3 5.6a.272.272 0 0 0-.05.15A.25.25 0 0 0 .5 6H1v3.745a.255.255 0 0 0 .255.255h2.49A.255.255 0 0 0 4 9.745V5.467a.253.253 0 0 1 .064-.167L5 4.5L3.2 2.316a.255.255 0 0 0-.423-.016zM3 7H2V6h1zm6.75-4h-.5a.25.25 0 0 0-.25.25V5.2L7.658 4.126a.253.253 0 0 0-.316 0L5.1 5.926a.253.253 0 0 0-.095.2v3.621a.253.253 0 0 0 .248.253h1.494A.253.253 0 0 0 7 9.747V8h1v1.747a.253.253 0 0 0 .253.253h1.494A.253.253 0 0 0 10 9.747V3.25A.25.25 0 0 0 9.75 3zM7 7H6V6h1zm2 0H8V6h1z" fill="currentColor"/>`),
		g.Group(children),
	)
}