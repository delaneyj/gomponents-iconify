package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TourOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14v7q0 .425-.288.713T6 22q-.425 0-.713-.288T5 21V3q0-.425.288-.713T6 2q.425 0 .713.288T7 3v1h12.525q.525 0 .825.438t.1.937L19 9l1.45 3.625q.2.5-.1.938t-.825.437H7Zm0-8v6v-6Zm5.5 5q.825 0 1.413-.588T14.5 9q0-.825-.588-1.413T12.5 7q-.825 0-1.413.588T10.5 9q0 .825.588 1.413T12.5 11ZM7 12h11.05l-1.2-3l1.2-3H7v6Z"/>`),
		g.Group(children),
	)
}