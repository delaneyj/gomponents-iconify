package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeTwoAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 14v-2h2v2H5Zm-2-2v-2h2v2H3Zm9-7V3h2v2h-2ZM4.5 7.5h3v-3h-3v3ZM3 8V4q0-.425.288-.713T4 3h4q.425 0 .713.288T9 4v4q0 .425-.288.713T8 9H4q-.425 0-.713-.288T3 8Zm1.5 11.5h3v-3h-3v3ZM3 20v-4q0-.425.288-.713T4 15h4q.425 0 .713.288T9 16v4q0 .425-.288.713T8 21H4q-.425 0-.713-.288T3 20ZM16.5 7.5h3v-3h-3v3ZM15 8V4q0-.425.288-.713T16 3h4q.425 0 .713.288T21 4v4q0 .425-.288.713T20 9h-4q-.425 0-.713-.288T15 8Zm-6 6v-2H7v-2h4v4H9Zm1-5V5h2v2h2v2h-4ZM5.25 6.75v-1.5h1.5v1.5h-1.5Zm0 12v-1.5h1.5v1.5h-1.5Zm12-12v-1.5h1.5v1.5h-1.5ZM16 18h-2q-.425 0-.713-.288T13 17q0-.425.288-.713T14 16h2v-2q0-.425.288-.713T17 13q.425 0 .713.288T18 14v2h2q.425 0 .713.288T21 17q0 .425-.288.713T20 18h-2v2q0 .425-.288.713T17 21q-.425 0-.713-.288T16 20v-2Z"/>`),
		g.Group(children),
	)
}