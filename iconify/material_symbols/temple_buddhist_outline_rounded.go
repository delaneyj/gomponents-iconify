package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleBuddhistOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-7.15q-1.3-.325-2.15-1.375T1 9.025q0-.425.288-.712T2 8.025q.425 0 .713.288T3 9.024q0 .8.588 1.388T4.974 11H6V8.85q-1.3-.325-2.15-1.375T3 5.025q0-.425.288-.713T4 4.025q.425 0 .713.288T5 5.024q0 .8.588 1.388T6.975 7H7.5l2.9-3.875q.3-.4.725-.6t.875-.2q.45 0 .875.2t.725.6L16.5 7h.525q.8 0 1.388-.588T19 5.025q0-.425.288-.713T20 4.026q.425 0 .713.288t.287.712q0 1.4-.85 2.45T18 8.85V11h1.025q.8 0 1.388-.587T21 9.025q0-.425.288-.713T22 8.025q.425 0 .713.288t.287.712q0 1.4-.85 2.45T20 12.85V20q0 .825-.588 1.413T18 22h-4q-.425 0-.713-.288T13 21v-3q0-.425-.288-.713T12 17q-.425 0-.713.288T11 18v3q0 .425-.288.713T10 22H6q-.825 0-1.413-.588T4 20Zm6-13h4l-2-2.675L10 7Zm-2 4h8V9H8v2Zm-2 9h3v-2q0-1.25.875-2.125T12 15q1.25 0 2.125.875T15 18v2h3v-7H6v7Zm6-7Zm0-6Zm0 4Z"/>`),
		g.Group(children),
	)
}