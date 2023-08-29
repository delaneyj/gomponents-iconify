package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RangerStationEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.334 0L5 1v3L1 6.055V10h3V7h3v3h3V6.055L6 4V2.25L6.334 2l1.332 1L9 2V0L7.666 1L6.334 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}