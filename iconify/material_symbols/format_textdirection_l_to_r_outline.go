package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextdirectionLToROutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15v-5q-1.65 0-2.825-1.175T5 6q0-1.65 1.175-2.825T9 2h8v2h-2v11h-2V4h-2v11H9Zm0-7V4q-.825 0-1.413.588T7 6q0 .825.588 1.413T9 8Zm0-2Zm8 16l-1.4-1.4l1.6-1.6H3v-2h14.2l-1.6-1.6L17 14l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}