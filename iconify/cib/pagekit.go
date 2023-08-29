package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pagekit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.203 0v32H16v-4.703H7.906V4.703h16.193v17.891h-8.094v4.703h12.797V0z"/>`),
		g.Group(children),
	)
}