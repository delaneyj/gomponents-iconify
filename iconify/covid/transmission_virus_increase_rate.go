package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionVirusIncreaseRate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 19.571a4.571 4.571 0 1 0 0-9.142a4.571 4.571 0 0 0 0 9.142ZM8.238 7h1.524M9 7v3.429m5.118-1.625l1.078 1.078m-.539-.539l-2.425 2.425M17 14.238v1.524M17 15h-3.429m1.625 5.118l-1.078 1.078m.539-.539l-2.425-2.425M9.762 23H8.238M9 23v-3.429m-5.118 1.625l-1.078-1.078m.539.539l2.425-2.425M1 15.762v-1.524M1 15h3.429M2.804 9.882l1.078-1.078m-.539.539l2.425 2.425M17 4l3-3l3 3m-3-3v7m0 4v3m0 4v3"/>`),
		g.Group(children),
	)
}