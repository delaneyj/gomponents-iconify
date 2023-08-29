package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12.896V21H2v-8.104l2.196-4.182H6.2V3h11.6v5.714h2.004L22 12.896Zm-6.2-4.182V5H8.2v3.714h7.6ZM20 13.39l-1.404-2.675H5.404L4 13.39V19h16v-5.61Z"/>`),
		g.Group(children),
	)
}