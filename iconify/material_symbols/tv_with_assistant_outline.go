package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvWithAssistantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 12.5q.425 0 .713-.288t.287-.712q0-.425-.288-.713T10.5 10.5q-.425 0-.713.288T9.5 11.5q0 .425.288.713t.712.287Zm3 0q.425 0 .713-.288t.287-.712q0-.425-.288-.713T13.5 10.5q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287Zm-6 0q.425 0 .713-.288T8.5 11.5q0-.425-.288-.713T7.5 10.5q-.425 0-.713.288T6.5 11.5q0 .425.288.713t.712.287Zm9 0q.425 0 .713-.288t.287-.712q0-.425-.288-.713T16.5 10.5q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287ZM4 21v-2q-.825 0-1.413-.588T2 17V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v11q0 .825-.588 1.413T20 19v2h-1l-.65-2H5.675L5 21H4Zm0-4h16V6H4v11Zm8-5.5Z"/>`),
		g.Group(children),
	)
}