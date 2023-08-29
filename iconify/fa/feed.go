package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M384 1216q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136zm512 123q2 28-17 48q-18 21-47 21H697q-25 0-43-16.5t-20-41.5q-22-229-184.5-391.5T58 774q-25-2-41.5-20T0 711V576q0-29 21-47q17-17 43-17h5q160 13 306 80.5T634 774q114 113 181.5 259t80.5 306zm512 2q2 27-18 47q-18 20-46 20h-143q-26 0-44.5-17.5T1137 1348q-12-215-101-408.5t-231.5-336t-336-231.5T60 270q-25-1-42.5-19.5T0 207V64q0-28 20-46Q38 0 64 0h3q262 13 501.5 120T994 414q187 186 294 425.5t120 501.5z"/>`),
		g.Group(children),
	)
}