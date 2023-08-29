package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOnlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16v-3q.425 0 .713-.288T9 12q0-.425-.288-.713T8 11V8h8v3q-.425 0-.713.288T15 12q0 .425.288.713T16 13v3H8Zm4-1.5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T12 13.5q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Zm0-2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T12 11.5q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Zm0-2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T12 9.5q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15ZM5 23V1h14v22H5Zm2-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}