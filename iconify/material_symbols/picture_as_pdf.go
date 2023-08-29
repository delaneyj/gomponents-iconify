package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureAsPdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12.5h1v-2h1q.425 0 .713-.288T12 9.5v-1q0-.425-.288-.713T11 7.5H9v5Zm1-3v-1h1v1h-1Zm3 3h2q.425 0 .713-.288T16 11.5v-3q0-.425-.288-.713T15 7.5h-2v5Zm1-1v-3h1v3h-1Zm3 1h1v-2h1v-1h-1v-1h1v-1h-2v5ZM8 18q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V6h2v14h14v2H4Z"/>`),
		g.Group(children),
	)
}