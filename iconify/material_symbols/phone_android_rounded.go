package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneAndroidRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 20h3q.2 0 .35-.15t.15-.35q0-.2-.15-.35T13.5 19h-3q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-7h10V6H7v10Z"/>`),
		g.Group(children),
	)
}