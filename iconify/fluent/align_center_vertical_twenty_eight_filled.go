package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.25 26a.75.75 0 0 1-.75-.75V23H10a2.75 2.75 0 0 1-2.75-2.75v-2.5A2.75 2.75 0 0 1 10 15h3.5v-2H7.25a2.75 2.75 0 0 1-2.75-2.75v-2.5A2.75 2.75 0 0 1 7.25 5h6.25V2.75a.75.75 0 0 1 1.5 0V5h5.75a2.75 2.75 0 0 1 2.75 2.75v2.5A2.75 2.75 0 0 1 20.75 13H15v2h3a2.75 2.75 0 0 1 2.75 2.75v2.5A2.75 2.75 0 0 1 18 23h-3v2.25a.75.75 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}