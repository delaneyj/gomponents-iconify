package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLaundryServiceSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v20H4Zm8-3q2.075 0 3.538-1.463T17 14q0-2.075-1.463-3.538T12 9q-2.075 0-3.538 1.463T7 14q0 2.075 1.463 3.538T12 19Zm0-1.7q-.675 0-1.288-.25t-1.062-.7l4.7-4.7q.45.45.7 1.063T15.3 14q0 1.35-.975 2.325T12 17.3ZM8 7q.425 0 .713-.288T9 6q0-.425-.288-.713T8 5q-.425 0-.713.288T7 6q0 .425.288.713T8 7Zm3 0q.425 0 .713-.288T12 6q0-.425-.288-.713T11 5q-.425 0-.713.288T10 6q0 .425.288.713T11 7Z"/>`),
		g.Group(children),
	)
}