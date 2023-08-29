package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExamBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H40a20 20 0 0 0-20 20v160a12 12 0 0 0 17.37 10.73L64 213.42l26.63 13.31a12 12 0 0 0 10.74 0L128 213.42l26.63 13.31a12 12 0 0 0 10.74 0L192 213.42l26.63 13.31A12 12 0 0 0 236 216V56a20 20 0 0 0-20-20Zm-4 160.58l-14.63-7.31a12 12 0 0 0-10.74 0L160 202.58l-26.63-13.31a12 12 0 0 0-10.74 0L96 202.58l-26.63-13.31a12 12 0 0 0-10.74 0L44 196.58V60h168ZM63.19 171A12 12 0 0 0 79 164.81l2.1-4.81h29.8l2.11 4.81a12 12 0 0 0 22-9.62l-28-64a12 12 0 0 0-22 0l-28 64A12 12 0 0 0 63.19 171Zm37.21-35h-8.8l4.4-10.06Zm35.6-8a12 12 0 0 1 12-12h8v-8a12 12 0 0 1 24 0v8h8a12 12 0 0 1 0 24h-8v8a12 12 0 0 1-24 0v-8h-8a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}