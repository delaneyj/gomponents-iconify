package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.06 3h19.88l1.124 18H16.28l-1-3H8.721l-1 3H.936L2.06 3Zm1.88 2l-.876 14H6.28l1-3h9.442l1 3h3.215L20.06 5H3.939ZM9 7.5v2h2v2H9v2H7v-2H5v-2h2v-2h2Zm7 0h2.004v2.004H16V7.5Zm0 3.996h2.004V13.5H16v-2.004Z"/>`),
		g.Group(children),
	)
}