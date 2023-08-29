package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteStackAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20V8.975q0-.825.6-1.4T9.025 7H20q.825 0 1.413.587T22 9v7.175q0 .4-.15.763t-.425.637l-3.85 3.85q-.275.275-.638.425t-.762.15H9q-.825 0-1.413-.588T7 20ZM2.025 6.25q-.15-.825.325-1.488t1.3-.812L14.5 2.025q.825-.15 1.488.325t.812 1.3L17.05 5H9Q7.35 5 6.175 6.175T5 9v9.55q-.4-.225-.688-.6t-.362-.85L2.025 6.25ZM13.5 15.5v2q0 .425.288.713t.712.287q.425 0 .713-.288t.287-.712v-2h2q.425 0 .713-.288t.287-.712q0-.425-.288-.713T17.5 13.5h-2v-2q0-.425-.288-.713T14.5 10.5q-.425 0-.713.288t-.287.712v2h-2q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287h2Z"/>`),
		g.Group(children),
	)
}