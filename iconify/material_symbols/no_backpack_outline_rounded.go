package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoBackpackOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 17.175l-2-2V8q0-.825-.588-1.413T16 6H8.825L7 4.15V3q0-.425.288-.713T8 2h1q.425 0 .713.288T10 3v1h4V3q0-.425.288-.713T15 2h1q.425 0 .713.288T17 3v1.15q1.3.35 2.15 1.4T20 8v9.175Zm-3.5-3.5L14.825 12h.675q.425 0 .713.288T16.5 13v.675ZM11.175 12v2H8.5q-.425 0-.713-.288T7.5 13q0-.425.288-.713T8.5 12h2.675Zm2.375-1.275Zm-2 3.65Zm-6.375-9.2l1.4 1.4q-.275.275-.425.638T6 8v12h12v-1.975L20 20q0 .825-.588 1.413T18 22H6q-.825 0-1.413-.588T4 20V8q0-.825.313-1.55t.862-1.275Zm13.9 16.725L2.1 4.925q-.275-.275-.275-.687t.275-.713q.3-.3.713-.3t.712.3L20.475 20.5q.275.275.287.688t-.287.712q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}