package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBinance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 8l2 2l4-4l4 4l2-2l-6-6zm0 8l2-2l4 4l3.5-3.5l2 2L12 22zm14-6l2 2l-2 2l-2-2zM4 10l2 2l-2 2l-2-2zm8 0l2 2l-2 2l-2-2z"/>`),
		g.Group(children),
	)
}