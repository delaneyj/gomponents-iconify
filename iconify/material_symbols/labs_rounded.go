package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17V8q-.825 0-1.413-.588T5 6V4q0-.825.588-1.413T7 2h10q.825 0 1.413.588T19 4v2q0 .825-.588 1.413T17 8v9q0 2.075-1.463 3.538T12 22Zm0-2q.975 0 1.75-.563T14.825 18H13q-.425 0-.713-.288T12 17q0-.425.288-.713T13 16h2v-1h-2q-.425 0-.713-.288T12 14q0-.425.288-.713T13 13h2v-1h-2q-.425 0-.713-.288T12 11q0-.425.288-.713T13 10h2V8H9v9q0 1.25.875 2.125T12 20Z"/>`),
		g.Group(children),
	)
}