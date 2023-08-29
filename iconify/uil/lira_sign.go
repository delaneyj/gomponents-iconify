package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiraSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 12a1 1 0 0 0-1 1a7.008 7.008 0 0 1-7 7v-8.865l5.217-1.159a1 1 0 0 0-.434-1.952L10 9.087V7.135l5.217-1.159a1 1 0 1 0-.434-1.952L10 5.087V3a1 1 0 0 0-2 0v2.531l-2.217.493a1 1 0 1 0 .434 1.952L8 7.58v1.95l-2.217.493a1 1 0 1 0 .434 1.952L8 11.58V21a1 1 0 0 0 1 1h1a9.01 9.01 0 0 0 9-9a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}