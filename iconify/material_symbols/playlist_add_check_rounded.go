package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistAddCheckRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h6q.425 0 .713.288T11 15q0 .425-.288.713T10 16H4Zm0-4q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h10q.425 0 .713.288T15 11q0 .425-.288.713T14 12H4Zm0-4q-.425 0-.713-.288T3 7q0-.425.288-.713T4 6h10q.425 0 .713.288T15 7q0 .425-.288.713T14 8H4Zm12.35 10.575q-.2 0-.375-.062t-.325-.213l-2.15-2.15q-.275-.275-.287-.687t.287-.713q.275-.275.688-.288t.712.263l1.45 1.425l3.525-3.525q.3-.3.713-.287t.712.312q.275.3.288.7t-.288.7l-4.25 4.25q-.15.15-.325.213t-.375.062Z"/>`),
		g.Group(children),
	)
}