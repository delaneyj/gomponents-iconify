package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLaundryServiceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Zm0-2h12V4H6v16Zm6-1q2.075 0 3.538-1.463T17 14q0-2.075-1.463-3.538T12 9q-2.075 0-3.538 1.463T7 14q0 2.075 1.463 3.538T12 19Zm0-1.7q-.65 0-1.263-.238T9.65 16.35l4.7-4.7q.475.475.713 1.088T15.3 14q0 1.375-.963 2.337T12 17.3ZM8 7q.425 0 .713-.288T9 6q0-.425-.288-.713T8 5q-.425 0-.713.288T7 6q0 .425.288.713T8 7Zm3 0q.425 0 .713-.288T12 6q0-.425-.288-.713T11 5q-.425 0-.713.288T10 6q0 .425.288.713T11 7ZM6 20V4v16Z"/>`),
		g.Group(children),
	)
}