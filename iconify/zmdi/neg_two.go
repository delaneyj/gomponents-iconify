package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NegTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M278 284h127v36H221v-32l89-97q10-11 19-22q7-8 12-18q4-7 6-14q2-8 2-14q0-9-3-18q-3-8-8-13q-5-7-12.5-10T308 79q-12 0-20 4q-9 4-15 10q-6 8-8 16q-3 9-3 19h-46q1-17 6-32q6-16 18-28t29-19q18-6 40-6q20 0 36 5q17 6 27 15q11 10 17 24t6 31q0 13-4 25q-5 12-12 25q-8 13-17 25q-13 15-23 25zM0 171h171v42H0v-42z"/>`),
		g.Group(children),
	)
}