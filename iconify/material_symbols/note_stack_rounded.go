package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteStackRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20V8.975q0-.825.6-1.4T9.025 7H20q.825 0 1.413.587T22 9v8l-5 5H9q-.825 0-1.413-.588T7 20ZM2.025 6.25q-.15-.825.325-1.488t1.3-.812L14.5 2.025q.825-.15 1.488.325t.812 1.3L17.05 5H9Q7.35 5 6.175 6.175T5 9v9.55q-.4-.225-.688-.6t-.362-.85L2.025 6.25ZM20 16h-3q-.425 0-.713.288T16 17v3l4-4Z"/>`),
		g.Group(children),
	)
}