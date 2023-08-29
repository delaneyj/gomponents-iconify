package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftwareDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11 5a1 1 0 1 1 2 0v7.158l3.243-3.243l1.414 1.414L12 15.986L6.343 10.33l1.414-1.414L11 12.158V5Z"/><path d="M4 14h2v4h12v-4h2v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4Z"/></g>`),
		g.Group(children),
	)
}