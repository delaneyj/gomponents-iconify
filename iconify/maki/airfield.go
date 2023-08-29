package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airfield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.818.682H4.773C4.09.682 4.09 0 4.773 0h5.454c.682 0 .682.682 0 .682H8.182S9 1.272 9 2.636V4h6v2L9 8l-.5 5l2.5 1.318V15H4v-.682L6.5 13L6 8L0 6V4h6V2.636c0-1.363.818-1.954.818-1.954z"/>`),
		g.Group(children),
	)
}