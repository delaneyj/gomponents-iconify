package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileVibrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 1h14v22H5V1Zm2 2v18h10V3H7ZM4 5v14H2V5h2Zm18 0v14h-2V5h2ZM11 17h2.004v2.004H11V17Z"/>`),
		g.Group(children),
	)
}