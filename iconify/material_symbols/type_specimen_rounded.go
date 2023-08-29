package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypeSpecimenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.2 12.2h3.65l.625 1.825q.075.2.263.338t.412.137q.375 0 .588-.313t.087-.662l-2.85-7.55q-.075-.225-.275-.35t-.425-.125h-.55q-.225 0-.425.125t-.275.35l-2.85 7.525q-.125.35.088.675t.612.325q.25 0 .438-.138t.262-.362l.625-1.8Zm.45-1.3l1.3-3.75h.1l1.3 3.75h-2.7ZM8 18q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4Z"/>`),
		g.Group(children),
	)
}