package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtTrackRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19q-.825 0-1.413-.588T1 17V7q0-.825.588-1.413T3 5h10q.825 0 1.413.588T15 7v10q0 .825-.588 1.413T13 19H3Zm4.5-5l-1-1.325q-.15-.2-.4-.188t-.4.213l-1.125 1.5q-.2.25-.038.525T5 15h6q.3 0 .45-.275t-.05-.525l-1.6-2.175q-.15-.2-.4-.2t-.4.2L7.5 14ZM18 19q-.425 0-.713-.288T17 18V6q0-.425.288-.713T18 5q.425 0 .713.288T19 6v12q0 .425-.288.713T18 19Zm4 0q-.425 0-.713-.288T21 18V6q0-.425.288-.713T22 5q.425 0 .713.288T23 6v12q0 .425-.288.713T22 19Z"/>`),
		g.Group(children),
	)
}