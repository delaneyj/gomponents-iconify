package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Raven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 18l-1.65 4l-1.85-.75l1.45-3.525q-2.65-.7-4.3-2.85T2 10V6q0-1.65 1.175-2.825T6 2q.55 0 1.05.187t1 .388L14 5l-4 1.5V8l7.85 5H10q-.825 0-1.413-.588T8 11H6q0 1.65 1.175 2.825T10 15h11l1 5h-2l-1-2h-5v4h-2v-4h-2ZM6 5q-.425 0-.713.288T5 6q0 .425.288.713T6 7q.425 0 .713-.288T7 6q0-.425-.288-.713T6 5Z"/>`),
		g.Group(children),
	)
}