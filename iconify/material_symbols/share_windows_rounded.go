package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareWindowsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 15q-.425 0-.713-.288T7 14V9q0-.825.588-1.413T9 7h8.15l-1.875-1.875q-.3-.3-.3-.713t.3-.712q.3-.3.712-.3t.713.3l3.6 3.6q.3.3.3.7t-.3.7l-3.6 3.6q-.275.275-.688.288T15.3 12.3q-.3-.3-.313-.713t.288-.712L17.15 9H9v5q0 .425-.288.713T8 15Zm-3 6q-.825 0-1.413-.588T3 19V5q0-.425.288-.713T4 4q.425 0 .713.288T5 5v14h12v-3q0-.425.288-.713T18 15q.425 0 .713.288T19 16v3q0 .825-.588 1.413T17 21H5Z"/>`),
		g.Group(children),
	)
}