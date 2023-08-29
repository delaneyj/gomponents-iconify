package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveHundredPx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M281 247q-12 13-12 14q-14 13-21 17q-26 19-59 12q-30-6-45-36q-1-2-2-3q-8 12-8 13q-21 28-57 29q-23 1-40-6q-36-16-37-56h40l1 4q5 20 21 24q21 5 35-9q9-10 10.5-25t-6.5-26q-8-12-23-14.5T52 191q-4 3-6 5q-3 5-11 5H6q9-53 20-110h111v33H58q-3 0-4 3q-1 4-7 39v3q21-19 52-14q27 5 42 34q1-1 1-3q2-2 2-3q21-43 69-36q26 3 46 22q1 1 23 24l2-2q22-23 24-25q17-16 38-19q23-3 43 5.5t30 30.5q16 38-2 74q-16 33-55 35q-29 1-51-17q-3-2-25-22q0-1-5-6zm-76 5q12 0 24-7q8-4 28-23q1-2 0-4q-2-1-9.5-8T236 199q-9-8-21-12q-14-4-25 2t-15 20q-1 4-1 8q-1 16 8 25.5t23 9.5zm99-32q22 20 24 21q13 12 30 11q25 0 30-24q1-7 0-15q-2-13-11.5-21t-22.5-6q-15 1-28 13q-1 1-22 21z"/>`),
		g.Group(children),
	)
}