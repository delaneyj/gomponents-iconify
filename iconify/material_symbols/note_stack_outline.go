package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteStackOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20V8.975q0-.825.6-1.4T9.025 7H20q.825 0 1.413.587T22 9v8l-5 5H9q-.825 0-1.413-.588T7 20ZM2.025 6.25q-.15-.825.325-1.488t1.3-.812L14.5 2.025q.825-.15 1.488.325t.812 1.3L17.05 5H15l-.175-1L4 5.925l1 5.65v6.975q-.4-.225-.687-.6t-.363-.85L2.025 6.25ZM9 9v11h7v-4h4V9H9Zm5.5 5.5Z"/>`),
		g.Group(children),
	)
}