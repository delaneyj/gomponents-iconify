package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zodiacaquarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M905 352L761 161L564 359l-.5.5l-.5.5l-7 6l-1 1l-1.5 1.5l-3.5 3.5l-5 4q-4 2-10 4q-10 4-21 4h-5q-8 0-16-3q-7-2-15-7q-4-3-5-3.5t-3.5-2t-3.5-2.5q-10-10-15-26l-89-177l-253 203q-19 19-45.5 19t-45-19T0 320.5T19 276L340 18q19-18 45-18t45 18q3 3 9 14l1 1l1 1q3 8 4 8l85 170L723 18q19-18 45.5-18T814 18q3 3 9 14v.5l1 .5l191 255q14 23 7 48.5t-30 39t-48.5 6.5t-38.5-30zM340 530q19-19 45-19t45 19q3 3 9 14h.5l.5 1l1 1q3 7 4 8l85 170l194-194q18-19 44.5-19t45.5 19q3 3 9 14l1 1l191 254q14 23 7 49t-30 39t-48.5 6.5T905 864L761 673L564 870l-1 1l-7 6q0 1-1 1.5t-1.5 1L550 883l-5 4q-4 2-10 4q-10 4-21 4q-3 1-5 1q-7-1-16-3q-7-3-15-8q-4-2-5-3t-3.5-2.5t-3.5-2.5q-10-9-15-25l-89-178l-253 203q-19 19-45.5 19t-45-18.5t-18.5-45T19 787z"/>`),
		g.Group(children),
	)
}