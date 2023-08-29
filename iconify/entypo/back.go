package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Back(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 7v6c0 1.103-.896 2-2 2H3v-3h13V8H5v2L1 6.5L5 3v2h12a2 2 0 0 1 2 2z"/>`),
		g.Group(children),
	)
}