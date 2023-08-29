package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 9h-3v3h-2V9H8V7h3V4h2v3h3v2zm-9.5 6a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5z"/>`),
		g.Group(children),
	)
}