package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M363 256q-28 0-46 18.5T299 320t18 45.5t46 18.5q27 0 45.5-18.5T427 320t-18.5-45.5T363 256zm0 85q-22 0-22-21t22-21q21 0 21 21t-21 21zm134-40l11-41q5-16-3.5-31.5T480 209l-98-32l-34-102q-12-30-41-30H164q-13-2-24.5 5T124 70L85 177l-59 30q-28 18-22 49l11 43q-15 7-15 21q0 21 21 21h47q15 43 60 43q27 0 45.5-18.5T192 320t-18.5-45.5T128 256q-45 0-60 43h-8l-13-54l60-27q1 0 6-5h250q1 0 2 1t2 1l98 32l-13 52h-25v42h64q21 0 21-21q0-16-15-19zm-369-2q21 0 21 21t-21 21t-21-21t21-21zm85-128h-79l30-86h49v86zm43 0V85h49l28 86h-77zm-64 128v42h107v-42H192z"/>`),
		g.Group(children),
	)
}