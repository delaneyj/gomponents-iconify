package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessMessages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20q-.825 0-1.413-.588T5 18V8.75L1 4h19q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H7Zm6.5-2.5l1.4-1.4l-1.1-1.1H18v-2H9l4.5 4.5ZM9 11h9l-4.5-4.5l-1.4 1.4L13.2 9H9v2Z"/>`),
		g.Group(children),
	)
}