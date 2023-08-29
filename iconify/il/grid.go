package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M301 8q10 0 17 7t6 16v278q0 10-6 17t-17 7H23q-10 0-16-7t-7-17V31q0-9 7-16t16-7h278zM0 448q0-9 7-16t16-7h278q10 0 17 7t6 16v278q0 10-6 17t-17 7H23q-10 0-16-7t-7-17V448zM417 31q0-9 7-16t16-7h278q10 0 17 7t6 16v278q0 10-6 17t-17 7H440q-10 0-16-7t-7-17V31zm23 719q-10 0-16-7t-7-17V448q0-9 7-16t16-7h278q10 0 17 7t6 16v278q0 10-6 17t-17 7H440z"/>`),
		g.Group(children),
	)
}