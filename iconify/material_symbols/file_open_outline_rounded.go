package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOpenOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V13q0 .425-.288.713T19 14q-.425 0-.713-.288T18 13V9h-4q-.425 0-.713-.288T13 8V4H6v16h8q.425 0 .713.288T15 21q0 .425-.288.713T14 22H6Zm13-2.575v1.225q0 .425-.288.713T18 21.65q-.425 0-.713-.287T17 20.65V17q0-.425.288-.713T18 16h3.65q.425 0 .713.288t.287.712q0 .425-.287.713T21.65 18H20.4l2.25 2.25q.275.275.275.688t-.275.712q-.3.3-.712.3t-.713-.3L19 19.425ZM6 20V4v16Z"/>`),
		g.Group(children),
	)
}