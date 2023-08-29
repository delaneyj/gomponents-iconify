package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrOffSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.4 10.6l-8-8.05q.65-.3 1.3-.425T12 2q2.5 0 4.25 1.75T18 8q0 .65-.15 1.3t-.45 1.3Zm-1.15 4.45L14.6 13.4q-.575.3-1.262.45T12 14q-2.5 0-4.25-1.75T6 8q0-.725.15-1.375T6.6 5.4L4.95 3.75l1.4-1.4l11.3 11.3l-1.4 1.4ZM20.5 22v-2h-2v-1.5h2v-2H22v2h2V20h-2v2h-1.5ZM13 22v-6h3.5q.6 0 1.05.45T18 17.5v1q0 .575-.263.888t-.637.512L18 22h-1.5l-.9-2h-1.1v2H13Zm1.5-3.5h2v-1h-2v1ZM0 22v-6h1.5v2h2v-2H5v6H3.5v-2.5h-2V22H0Zm6.5 0v-6H10q.6 0 1.05.45t.45 1.05v3q0 .6-.45 1.05T10 22H6.5ZM8 20.5h2v-3H8v3ZM20.5 22v-2h-2v-1.5h2v-2H22v2h2V20h-2v2h-1.5ZM13 22v-6h3.5q.6 0 1.05.45T18 17.5v1q0 .575-.263.888t-.637.512L18 22h-1.5l-.9-2h-1.1v2H13Zm1.5-3.5h2v-1h-2v1ZM0 22v-6h1.5v2h2v-2H5v6H3.5v-2.5h-2V22H0Zm6.5 0v-6H10q.6 0 1.05.45t.45 1.05v3q0 .6-.45 1.05T10 22H6.5ZM8 20.5h2v-3H8v3Z"/>`),
		g.Group(children),
	)
}