package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationCircleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 56C145.72 56 56 145.72 56 256s89.72 200 200 200s200-89.72 200-200S366.28 56 256 56Zm0 82a26 26 0 1 1-26 26a26 26 0 0 1 26-26Zm64 226H200v-32h44v-88h-32v-32h64v120h44Z"/>`),
		g.Group(children),
	)
}