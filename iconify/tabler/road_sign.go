package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.446 2.6l7.955 7.954a2.045 2.045 0 0 1 0 2.892l-7.955 7.955a2.045 2.045 0 0 1-2.892 0l-7.955-7.955a2.045 2.045 0 0 1 0-2.892l7.955-7.955a2.045 2.045 0 0 1 2.892 0z"/><path d="M9 14v-2c0-.59.414-1 1-1h5"/><path d="m13 9l2 2l-2 2"/></g>`),
		g.Group(children),
	)
}