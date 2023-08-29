package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 0h1v11.5c0 1.381-1.567 2.5-3.5 2.5S9 12.881 9 11.5S10.567 9 12.5 9c.979 0 1.865.287 2.5.751V4L7 5.778V13.5C7 14.881 5.433 16 3.5 16S0 14.881 0 13.5S1.567 11 3.5 11c.979 0 1.865.287 2.5.751V2l9-2z"/>`),
		g.Group(children),
	)
}