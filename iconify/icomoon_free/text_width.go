package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 14v2l-3-2.5L4 11v2h8v-2l3 2.5l-3 2.5v-2zm9-13v4l-1-2H9v7h2v1H5v-1h2V3H4L3 5V1z"/>`),
		g.Group(children),
	)
}