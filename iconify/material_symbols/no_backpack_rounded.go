package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoBackpackRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.5 13.675l3.5 3.5V8q0-1.4-.85-2.45T17 4.15V3q0-.425-.288-.713T16 2h-1q-.425 0-.713.288T14 3v1h-4V3q0-.425-.288-.713T9 2H8q-.425 0-.713.288T7 3v1.15L14.825 12h.675q.425 0 .713.288T16.5 13v.675Zm2.575 8.225l-2.5-2.5h2.8L20 20q0 .825-.588 1.413T18 22H6q-.825 0-1.413-.588T4 20V8q0-.825.313-1.55t.862-1.275l.975.975v2.825L2.1 4.925q-.275-.275-.275-.687t.275-.713q.3-.3.713-.3t.712.3L20.475 20.5q.275.275.287.688t-.287.712q-.275.275-.7.275t-.7-.275Zm-7.9-7.9l-2-2H8.5q-.425 0-.713.288T7.5 13q0 .425.288.713T8.5 14h2.675Z"/>`),
		g.Group(children),
	)
}