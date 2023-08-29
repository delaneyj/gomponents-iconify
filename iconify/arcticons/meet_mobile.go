package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeetMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="35.291" cy="19.357" r="5.156" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.175 34.202a9.728 9.728 0 0 0 2.448.31a9.7 9.7 0 0 0 8.098-4.355a9.7 9.7 0 0 0 8.099 4.356c3.526 0 6.604-1.886 8.305-4.697c1.701 2.811 4.78 4.697 8.305 4.697c.71 0 1.403-.077 2.07-.221h0L16.547 13.487l-10.834 7.29a2.746 2.746 0 0 0 3.066 4.556l6.826-4.592l4.199 3.371l-13.629 10.09Z"/>`),
		g.Group(children),
	)
}