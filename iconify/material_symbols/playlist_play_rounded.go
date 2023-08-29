package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistPlayRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h6q.425 0 .713.288T11 15q0 .425-.288.713T10 16H4Zm0-4q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h10q.425 0 .713.288T15 11q0 .425-.288.713T14 12H4Zm0-4q-.425 0-.713-.288T3 7q0-.425.288-.713T4 6h10q.425 0 .713.288T15 7q0 .425-.288.713T14 8H4Zm12.775 12.475q-.125.075-.25.075t-.25-.05q-.125-.05-.2-.163T16 20.075v-6.15q0-.15.075-.263t.2-.162q.125-.05.25-.05t.25.075l4.6 3.05q.125.075.175.188t.05.237q0 .125-.05.238t-.175.187l-4.6 3.05Z"/>`),
		g.Group(children),
	)
}