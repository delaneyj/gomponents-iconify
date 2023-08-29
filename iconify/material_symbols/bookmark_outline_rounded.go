package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18l-4.2 1.8q-1 .425-1.9-.163T5 17.976V5q0-.825.588-1.413T7 3h10q.825 0 1.413.588T19 5v12.975q0 1.075-.9 1.663t-1.9.162L12 18Zm0-2.2l5 2.15V5H7v12.95l5-2.15ZM12 5H7h10h-5Z"/>`),
		g.Group(children),
	)
}