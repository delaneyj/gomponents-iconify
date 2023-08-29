package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioFileRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V20q0 .825-.588 1.413T18 22H6Zm7-14q0 .425.288.713T14 9h4l-5-5v4Zm-2.25 11q.95 0 1.6-.65t.65-1.6V13h2q.425 0 .713-.288T16 12q0-.425-.288-.713T15 11h-2q-.425 0-.713.288T12 12v2.875q-.275-.2-.588-.288t-.662-.087q-.95 0-1.6.65t-.65 1.6q0 .95.65 1.6t1.6.65Z"/>`),
		g.Group(children),
	)
}