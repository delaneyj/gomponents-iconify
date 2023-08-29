package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tapas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23v-9H4q-1.05 0-1.775-.725T1.5 11.5q0-1.05.725-1.775T4 9h2V8H4q-1.05 0-1.775-.725T1.5 5.5q0-1.05.725-1.775T4 3h2V1h2v2h2q1.05 0 1.775.725T12.5 5.5q0 1.05-.725 1.775T10 8H8v1h2q1.05 0 1.775.725T12.5 11.5q0 1.05-.725 1.775T10 14H8v9H6Zm9 0v-2h2v-7.15q-1.3-.35-2.15-1.4T14 10V1h8v9q0 1.4-.85 2.45T19 13.85V21h2v2h-6Zm1-17h4V3h-4v3Z"/>`),
		g.Group(children),
	)
}