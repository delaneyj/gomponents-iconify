package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.9 4.1v-2h8v8h-2V5.516l-5.779 5.778l-1.414-1.414l5.778-5.78H13.9Zm-9.8 9.8h-2v8h8v-2H5.516l5.778-5.779l-1.414-1.414l-5.78 5.778V13.9Z"/>`),
		g.Group(children),
	)
}