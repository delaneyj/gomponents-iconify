package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectLargeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V7h12q.825 0 1.413.588T17 9v12H5Zm1-3h8q.3 0 .45-.275t-.05-.525L12 14.025q-.15-.2-.4-.2t-.4.2L9 17l-1.2-1.625q-.15-.2-.4-.2t-.4.2L5.6 17.2q-.2.25-.063.525T6 18ZM3 5q0-.825.588-1.413T5 3v2H3Zm8 0V3h2v2h-2Zm8 0V3q.825 0 1.413.588T21 5h-2ZM7 5V3h2v2H7Zm12 8v-2h2v2h-2Zm0 8v-2h2q0 .825-.588 1.413T19 21Zm0-12V7h2v2h-2Zm0 8v-2h2v2h-2ZM15 5V3h2v2h-2Z"/>`),
		g.Group(children),
	)
}