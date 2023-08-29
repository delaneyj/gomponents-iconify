package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorLampTorchiereOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 4l-1.3 3h-1.4L10 4h4m3-2H7l3 7h4l3-7m-4 18h3v2H8v-2h3V10h2v10Z"/>`),
		g.Group(children),
	)
}