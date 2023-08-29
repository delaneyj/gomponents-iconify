package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wordpress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 232q0 95 67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2T69.5 69.5T2 232zm17 0q0-44 19-87l101 278q-54-26-87-77.5T19 232zm213 213q-32 0-60-9l64-185l65 179q1 0 1 1t1 2q-38 12-71 12zm172-217q16-40 16-76q0-15-1-22q26 49 26 102q0 57-29 106.5T339 416zm-46-65q18 32 18 58t-17 69l-21 71l-77-229l25-2q5 0 7-4.5t-.5-8.5t-8.5-4q-34 2-56 2l-56-2q-6 0-8.5 4t-.5 8.5t7 4.5q11 2 23 2l33 91l-47 140l-77-231l24-2q6 0 7.5-4.5t-1-8.5t-7.5-4q-35 2-57 2H54q29-44 76-70t102-26q81 0 144 56h-3q-15 0-25.5 11.5T337 113q0 16 21 50z"/>`),
		g.Group(children),
	)
}