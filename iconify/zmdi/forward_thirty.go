package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M119 272h9q6 0 10.5-4.5t4.5-8.5v-4q-2-2-2-4t-4-2h-11q-2 2-4.5 2t-2.5 4v4H98q0-6 2-10.5t6.5-8.5t8.5-4q1 0 5.5-1t5.5-1q8 0 13 2q2 1 5 2t5 2q3 1 7 9q2 4 2 10v7q-2 4-2 6q0 4-5 4q-2 0-6 5q9 4 11 8q4 9 4 13q0 8-2 11q-1 1-3 4t-4 4q-4 4-10 4q-2 0-6.5 1t-6.5 1q-9 0-11-2q-1-1-5-2t-5-2q-3-1-7-8q-2-5-2-13h17v4q2 2 2 4t5 2h10q2-2 4.5-2t2.5-4v-11q-2-2-2-4t-5-2h-13v-15zm122 15q0 13-2 17l-6 13q-7 6-11 6q-2 0-6.5 1t-6.5 1q-8 0-13-2q-2-1-5-3t-5-3q-3-1-7-13q-2-6-2-17v-15q0-13 2-17l7-13q6-6 10-6q2 0 6.5-1t6.5-1q9 0 13 2q2 1 5 3t6 3q2 1 6 13q2 6 2 17v15zm-19-17v-11q-2-4-2-6l-5-4q-2-3-6-3t-6 3l-5 4q-2 4-2 6v43q2 4 2 6l5 5q2 2 6 2t6-2l5-5q2-4 2-6v-32zM0 261.5Q0 191 50 141t121-50V5l106 107l-106 107v-86q-53 0-90.5 38T43 261.5t37.5 90t90 37.5t90.5-37.5t38-90.5h42q0 71-50 121t-120.5 50T50 382T0 261.5z"/>`),
		g.Group(children),
	)
}