package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestSunblock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7V4h2v3H2Zm0 13v-3h2v3H2ZM7.25 9.175L5.825 7.75L7.95 5.625L9.375 7.05L7.25 9.175Zm.7 9.2L5.825 16.25l1.425-1.425l2.125 2.125l-1.425 1.425ZM8 13v-2h6l-1.8-2.4l1.6-1.2l3 4q.375.5.1 1.05T16 13H8Zm12.5 6q-.625 0-1.063-.438T19 17.5v-11q0-.625.438-1.063T20.5 5q.625 0 1.063.438T22 6.5v11q0 .625-.438 1.063T20.5 19ZM2 16v-2q.825 0 1.413-.588T4 12q0-.825-.588-1.413T2 10V8q1.65 0 2.825 1.175T6 12q0 1.65-1.175 2.825T2 16Z"/>`),
		g.Group(children),
	)
}