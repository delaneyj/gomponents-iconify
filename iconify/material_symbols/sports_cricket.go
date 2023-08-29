package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsCricket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 14.2L12.2 17q-.3.3-.7.3t-.7-.3L2.3 8.5Q2 8.2 2 7.825t.3-.675l2.8-2.8q.3-.3.725-.3t.725.3L15 12.8q.3.3.3.7t-.3.7Zm3.6 7.8l-4.25-4.25l1.4-1.4L20 20.6L18.6 22Zm-.1-13q-1.45 0-2.475-1.025T15 5.5q0-1.45 1.025-2.475T18.5 2q1.45 0 2.475 1.025T22 5.5q0 1.45-1.025 2.475T18.5 9Z"/>`),
		g.Group(children),
	)
}