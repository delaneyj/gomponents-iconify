package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapasRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22v-8H4q-1.05 0-1.775-.725T1.5 11.5q0-1.05.725-1.775T4 9h2V8H4q-1.05 0-1.775-.725T1.5 5.5q0-1.05.725-1.775T4 3h2V2q0-.425.288-.713T7 1q.425 0 .713.288T8 2v1h2q1.05 0 1.775.725T12.5 5.5q0 1.05-.725 1.775T10 8H8v1h2q1.05 0 1.775.725T12.5 11.5q0 1.05-.725 1.775T10 14H8v8q0 .425-.288.713T7 23q-.425 0-.713-.288T6 22Zm11-1v-7.15q-1.3-.35-2.15-1.4T14 10V2q0-.425.288-.713T15 1h6q.425 0 .713.288T22 2v8q0 1.4-.85 2.45T19 13.85V21h1q.425 0 .713.288T21 22q0 .425-.288.713T20 23h-4q-.425 0-.713-.288T15 22q0-.425.288-.713T16 21h1ZM16 6h4V3h-4v3Z"/>`),
		g.Group(children),
	)
}