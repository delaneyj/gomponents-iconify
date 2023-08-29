package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.95 7.325L4.225 6.6q-.325-.3-.338-.7t.313-.7q.3-.3.712-.313t.713.288l.725.725q.275.275.288.688T6.35 7.3q-.275.275-.687.288t-.713-.263ZM12 5q-.425 0-.713-.288T11 4V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v1q0 .425-.288.713T12 5Zm5.7 2.325q-.3-.3-.3-.725t.3-.725l.7-.7q.275-.275.688-.275t.712.3q.275.275.275.7t-.275.7l-.7.7q-.3.3-.7.313t-.7-.288ZM11 22q-.825 0-1.413-.588T9 20v-3l-2.425-2.425q-.275-.275-.425-.638T6 13.175V10q0-.425.288-.713T7 9h10q.425 0 .713.288T18 10v3.175q0 .4-.15.763t-.425.637L15 17v3q0 .825-.588 1.413T13 22h-2Z"/>`),
		g.Group(children),
	)
}