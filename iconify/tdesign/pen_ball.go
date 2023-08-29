package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.155 1.374l5.471 5.472l-1.414 1.414l-5.471-5.472l1.414-1.414Zm-2.802 2.833l5.441 5.44L7.397 22.002H2v-5.397L14.353 4.207Zm.002 2.83L4 17.43v2.571h2.57L16.964 9.645l-2.608-2.607Zm7.853 5.83l-6.244 6.243l-1.414-1.414l6.244-6.244l1.414 1.414Z"/>`),
		g.Group(children),
	)
}