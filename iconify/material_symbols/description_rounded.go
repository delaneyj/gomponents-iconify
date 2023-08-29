package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DescriptionRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18h6q.425 0 .713-.288T16 17q0-.425-.288-.713T15 16H9q-.425 0-.713.288T8 17q0 .425.288.713T9 18Zm0-4h6q.425 0 .713-.288T16 13q0-.425-.288-.713T15 12H9q-.425 0-.713.288T8 13q0 .425.288.713T9 14Zm-3 8q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V20q0 .825-.588 1.413T18 22H6Zm7-14q0 .425.288.713T14 9h4l-5-5v4Z"/>`),
		g.Group(children),
	)
}