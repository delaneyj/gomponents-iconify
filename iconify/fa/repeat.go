package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1536 128v448q0 26-19 45t-45 19h-448q-42 0-59-40q-17-39 14-69l138-138Q969 256 768 256q-104 0-198.5 40.5T406 406T296.5 569.5T256 768t40.5 198.5T406 1130t163.5 109.5T768 1280q119 0 225-52t179-147q7-10 23-12q15 0 25 9l137 138q9 8 9.5 20.5t-7.5 22.5q-109 132-264 204.5T768 1536q-156 0-298-61t-245-164t-164-245T0 768t61-298t164-245T470 61T768 0q147 0 284.5 55.5T1297 212l130-129q29-31 70-14q39 17 39 59z"/>`),
		g.Group(children),
	)
}