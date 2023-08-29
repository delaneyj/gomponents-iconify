package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileBlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 1h16v4h-2V3H6v18h12v-2h2v4H4V1Zm14 7.5a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.483 3.483 0 0 0 18 8.5Zm3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745ZM12.5 12a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM11 17h2.004v2.004H11V17Z"/>`),
		g.Group(children),
	)
}