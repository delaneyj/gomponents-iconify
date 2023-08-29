package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePresentRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q1.675 0 2.838-1.175T16 15v-3q0-.425-.288-.713T15 11q-.425 0-.713.288T14 12v3q0 .825-.575 1.413T12 17q-.825 0-1.413-.588T10 15V9.5q0-.225.15-.363T10.5 9q.225 0 .363.138T11 9.5V14q0 .425.288.713T12 15q.425 0 .713-.288T13 14V9.5q0-1.05-.725-1.775T10.5 7q-1.05 0-1.775.725T8 9.5V15q0 1.65 1.175 2.825T12 19Zm-6 3q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h9l5 5v13q0 .825-.588 1.413T18 22H6Zm8-18v3q0 .425.288.713T15 8h3l-4-4Z"/>`),
		g.Group(children),
	)
}