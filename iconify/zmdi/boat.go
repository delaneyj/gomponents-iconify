package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 432h43v43h-43q-44 0-85-21q-41 20-86 20t-85-20q-42 21-85 21H0v-43h43q45 0 85-28q39 27 85.5 27t85.5-27q40 28 85 28zM42 389L1 247q-3-8 1-17q4-8 13-10l28-9v-99q0-18 12.5-30.5T85 69h64V5h128v64h64q18 0 30.5 12.5T384 112v99l27 9q9 2 13 10t1 17l-40 142h-1q-48 0-85-42q-38 42-86 42t-85-42q-37 42-85 42h-1zm43-277v85l128-42l128 42v-85H85z"/>`),
		g.Group(children),
	)
}