package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnRightLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.472 15.447a.75.75 0 0 1-.393-1.031l2.5-5.001a.75.75 0 1 1 1.342.67L6.299 13.33l9.818-4.32L5.455 4.44a.75.75 0 0 1 .59-1.378l11.192 4.796c1.005.431 1.012 1.853.011 2.293L6.782 14.756l3.246 1.298a.75.75 0 1 1-.556 1.392l-5-1.999Z"/>`),
		g.Group(children),
	)
}