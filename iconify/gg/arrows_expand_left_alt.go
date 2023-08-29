package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandLeftAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.1 2.1v2H5.516l5.778 5.779l-1.414 1.414L4.1 5.515V10.1h-2v-8h8Zm11.8 11.8h-2v4.585l-5.779-5.778l-1.414 1.414l5.778 5.778H13.9v2h8v-8Zm-5.657-4.728l-1.415-1.415l-7.07 7.072l1.414 1.414l7.07-7.071Z"/>`),
		g.Group(children),
	)
}