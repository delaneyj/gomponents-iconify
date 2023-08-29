package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCompactAltOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 16.5h2q.425 0 .713-.288t.287-.712v-2q0-.425-.288-.713T10.5 12.5h-2q-.425 0-.713.288T7.5 13.5v2q0 .425.288.713t.712.287Zm0-5h2q.425 0 .713-.288t.287-.712v-2q0-.425-.288-.713T10.5 7.5h-2q-.425 0-.713.288T7.5 8.5v2q0 .425.288.713t.712.287Zm5 5h2q.425 0 .713-.288t.287-.712v-2q0-.425-.288-.713T15.5 12.5h-2q-.425 0-.713.288t-.287.712v2q0 .425.288.713t.712.287Zm0-5h2q.425 0 .713-.288t.287-.712v-2q0-.425-.288-.713T15.5 7.5h-2q-.425 0-.713.288T12.5 8.5v2q0 .425.288.713t.712.287ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}