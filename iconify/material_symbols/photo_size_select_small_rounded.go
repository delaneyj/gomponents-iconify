package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectSmallRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-8h8q.825 0 1.413.588T13 13v8H5ZM3 9V7h2v2H3Zm0-4q0-.825.588-1.413T5 3v2H3Zm2 14h6q.3 0 .45-.275t-.05-.525l-1.6-2.175q-.125-.2-.388-.2t-.412.2L7.5 18l-1-1.325q-.15-.2-.4-.187t-.4.212l-1.125 1.5q-.2.25-.05.525T5 19ZM7 5V3h2v2H7Zm4 0V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2q0 .825-.588 1.413T19 21Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm0-4V3q.825 0 1.413.588T21 5h-2Z"/>`),
		g.Group(children),
	)
}