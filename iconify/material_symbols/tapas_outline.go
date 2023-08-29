package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapasOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23v-9H4q-1.05 0-1.775-.725T1.5 11.5q0-1.05.725-1.775T4 9h2V8H4q-1.05 0-1.775-.725T1.5 5.5q0-1.05.725-1.775T4 3h2V1h2v2h2q1.05 0 1.775.725T12.5 5.5q0 1.05-.725 1.775T10 8H8v1h2q1.05 0 1.775.725T12.5 11.5q0 1.05-.725 1.775T10 14H8v9H6ZM4 12h6q.2 0 .35-.15t.15-.35q0-.2-.15-.35T10 11H4q-.2 0-.35.15t-.15.35q0 .2.15.35T4 12Zm0-6h6q.2 0 .35-.15t.15-.35q0-.2-.15-.35T10 5H4q-.2 0-.35.15t-.15.35q0 .2.15.35T4 6Zm11 17v-2h2v-7.15q-1.3-.35-2.15-1.4T14 10V1h8v9q0 1.4-.85 2.45T19 13.85V21h2v2h-6Zm3-11q.825 0 1.413-.588T20 10V8h-4v2q0 .825.588 1.413T18 12Zm-2-6h4V3h-4v3ZM3.5 12v-1v1Zm0-6V5v1ZM18 8Z"/>`),
		g.Group(children),
	)
}