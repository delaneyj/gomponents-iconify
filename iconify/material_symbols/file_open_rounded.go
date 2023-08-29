package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOpenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16h3.65q.425 0 .713.288t.287.712q0 .425-.287.713T21.65 18H20.4l2.25 2.25q.275.275.275.688t-.275.712q-.3.3-.712.3t-.713-.3L19 19.425v1.225q0 .425-.288.713T18 21.65q-.425 0-.713-.287T17 20.65V17q0-.425.288-.713T18 16ZM13 4v4q0 .425.288.713T14 9h4l-5-5ZM6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V13q0 .425-.288.713T19 14h-3q-.425 0-.713.288T15 15v6q0 .425-.288.713T14 22H6q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2Z"/>`),
		g.Group(children),
	)
}