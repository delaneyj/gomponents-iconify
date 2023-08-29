package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 11v-1q0-1.8.813-3.288T8 4.276L6.125 2.4L7 1.5l2.125 2.125q.65-.3 1.388-.462T12 3q.75 0 1.488.163t1.387.462L17 1.5l.875.9L16 4.275q1.375.95 2.188 2.438T19 10v1H5Zm10-2q.425 0 .713-.288T16 8q0-.425-.288-.713T15 7q-.425 0-.713.288T14 8q0 .425.288.713T15 9ZM9 9q.425 0 .713-.288T10 8q0-.425-.288-.713T9 7q-.425 0-.713.288T8 8q0 .425.288.713T9 9Zm3 14q-2.925 0-4.963-2.038T5 16v-4h14v4q0 2.925-2.038 4.963T12 23Z"/>`),
		g.Group(children),
	)
}