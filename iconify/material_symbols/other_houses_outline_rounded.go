package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OtherHousesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-7.375L3 12.4q-.35.25-.75.2t-.65-.4q-.25-.35-.188-.75t.388-.65l8.975-6.875q.275-.2.588-.3t.637-.1q.325 0 .638.1t.587.3l9 6.875q.325.25.375.65t-.2.75q-.25.325-.65.375t-.725-.2L20 11.625V19q0 .825-.587 1.413T18 21H6q-.825 0-1.413-.588T4 19Zm2 0h12v-8.9l-6-4.575L6 10.1V19Zm2-4q.425 0 .713-.288T9 14q0-.425-.288-.713T8 13q-.425 0-.713.288T7 14q0 .425.288.713T8 15Zm4 0q.425 0 .713-.288T13 14q0-.425-.288-.713T12 13q-.425 0-.713.288T11 14q0 .425.288.713T12 15Zm4 0q.425 0 .713-.288T17 14q0-.425-.288-.713T16 13q-.425 0-.713.288T15 14q0 .425.288.713T16 15ZM6 19h12H6Z"/>`),
		g.Group(children),
	)
}