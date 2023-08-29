package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PitchEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4 2a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm6.5 6H9L8 5L7 3.25L8 3l2.3 1a.531.531 0 0 0 .36-1L8 2H6L4 3L3 4H1.47a.5.5 0 0 0 0 1H4l1-1l1 2l-2 1v3.5a.5.5 0 0 0 1 0V7.39L7 7l1 2h2.5a.5.5 0 0 0 0-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}