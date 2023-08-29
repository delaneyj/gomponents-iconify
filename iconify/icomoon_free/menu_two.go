package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 3h14v3H0V3zm0 4h14v3H0V7zm0 4h14v3H0v-3zm15.5-2l3 3l3-3zm6-1l-3-3l-3 3z"/>`),
		g.Group(children),
	)
}