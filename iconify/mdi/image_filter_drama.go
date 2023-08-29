package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFilterDrama(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20h12q1.88 0 3.19-1.31T23 15.5q0-1.7-1.21-3.1Q20.58 11 19 11q0-2.95-2-4.97Q14.95 4 12 4Q9.88 4 8.29 5.08Q6.7 6.15 6.25 8q2.55.45 4.15 2.19Q12 11.93 12 14.5h-2q0-1.82-1.34-3.16Q7.33 10 5.5 10t-3.16 1.34Q1 12.68 1 14.5q0 2.3 1.6 3.9T6.5 20Z"/>`),
		g.Group(children),
	)
}