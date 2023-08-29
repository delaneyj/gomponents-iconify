package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdFlashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M128 298l64 64v118h128V362l64-64V176l-255.2.4L128 298zM234.8 32h42.4v64h-42.4V32zM80 110.4L109.9 80l44.9 45.6-29.9 30.4L80 110.4zm277.1 15.2l45-45.5 29.9 30.4-44.9 45.5-30-30.4z" fill="currentColor"/>`),
		g.Group(children),
	)
}