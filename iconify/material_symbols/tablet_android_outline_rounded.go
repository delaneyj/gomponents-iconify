package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletAndroidOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23q-.825 0-1.413-.588T3 21V3q0-.825.588-1.413T5 1h14q.825 0 1.413.588T21 3v18q0 .825-.588 1.413T19 23H5Zm0-5v3h14v-3H5Zm5.5 2h3q.2 0 .35-.15t.15-.35q0-.2-.15-.35T13.5 19h-3q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15ZM5 16h14V6H5v10ZM5 4h14V3H5v1Zm0 0V3v1Zm0 14v3v-3Z"/>`),
		g.Group(children),
	)
}