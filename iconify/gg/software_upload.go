package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftwareUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11 14.986a1 1 0 1 0 2 0V7.828l3.243 3.243l1.414-1.414L12 4L6.343 9.657l1.414 1.414L11 7.83v7.157Z"/><path d="M4 14h2v4h12v-4h2v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4Z"/></g>`),
		g.Group(children),
	)
}