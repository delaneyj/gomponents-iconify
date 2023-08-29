package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reddit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 189q0 28-27 39q2 9 2 19q0 51-55.5 87.5t-134 36.5t-134-36.5T23 247q0-10 2-20q-25-11-25-38q0-18 12.5-30.5T42 146q19 0 32 15q52-36 129-39l35-104q3-7 10-5l83 20q1 0 3 1q8-20 30-20q13 0 23 10t10 23.5T387 71t-23 10q-14 0-23.5-9.5T331 48q-2 1-3 0l-77-18l-31 92q79 2 132 40q13-16 33-16q17 0 29.5 12.5T427 189zm-311 33.5q0 12.5 9 21.5t21.5 9t21.5-9t9-21.5t-9-22t-21.5-9.5t-21.5 9.5t-9 22zM282 307q4-3 .5-6.5t-7.5-.5q-18 19-62 19t-62-19q-3-3-6.5.5t-.5 6.5q21 22 70 22q47 0 68-22zm-1.5-54q12.5 0 21.5-9t9-21.5t-9-22t-21.5-9.5t-22 9.5t-9.5 22t9.5 21.5t22 9z"/>`),
		g.Group(children),
	)
}