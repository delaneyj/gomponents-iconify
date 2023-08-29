package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videocam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m299 160l85-85v234l-85-85v75q0 8-6.5 14.5T277 320H21q-8 0-14.5-6.5T0 299V85q0-8 6.5-14.5T21 64h256q9 0 15.5 6.5T299 85v75z"/>`),
		g.Group(children),
	)
}