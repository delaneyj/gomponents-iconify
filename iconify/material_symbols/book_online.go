package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOnline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16q-.425 0-.713-.288T8 15v-2q.425 0 .713-.288T9 12q0-.425-.288-.713T8 11V9q0-.425.288-.713T9 8h6q.425 0 .713.288T16 9v2q-.425 0-.713.288T15 12q0 .425.288.713T16 13v2q0 .425-.288.713T15 16H9Zm3-1.5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T12 13.5q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Zm0-2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T12 11.5q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Zm0-2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T12 9.5q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}