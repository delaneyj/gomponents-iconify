package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 20l-4.2 1.8q-1 .425-1.9-.163T3 19.976V7q0-.825.588-1.413T5 5h10q.825 0 1.413.588T17 7v12.975q0 1.075-.9 1.663t-1.9.162L10 20Zm0-2.2l5 2.15V7H5v12.95l5-2.15Zm9 2.2V3H6V1h13q.825 0 1.413.588T21 3v17h-2ZM10 7H5h10h-5Z"/>`),
		g.Group(children),
	)
}