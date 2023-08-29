package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MosqueSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 8.05V7.1q0-1 .475-1.85t1.275-1.4L12 1l4.25 2.85q.8.55 1.275 1.4T18 7.1v.95H6ZM1 21V8.725Q.55 8.45.275 8.012T0 7q0-.575.6-1.4T2 4q.8.775 1.4 1.6T4 7q0 .575-.275 1.012T3 8.725V13h2V9.05h14V13h2V8.725q-.45-.275-.725-.713T20 7q0-.575.6-1.4T22 4q.8.775 1.4 1.6T24 7q0 .575-.275 1.012T23 8.725V21h-9v-6h-4v6H1Z"/>`),
		g.Group(children),
	)
}