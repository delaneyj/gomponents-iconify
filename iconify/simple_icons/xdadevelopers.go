package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xdadevelopers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.84 3.052V0h7.843v17.583H13.84v-3.024h4.591V3.052zM5.569 14.53V3.024h4.592V0H2.318v17.583H6.98L10.16 24v-9.483z"/>`),
		g.Group(children),
	)
}