package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gastroenterology(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-6q0-1.25.875-2.125T6 13h2q1.25 0 2.125-.875T11 10q0-.425-.288-.713T10 9q-.825 0-1.413-.588T8 7V2h6v2q2.925 0 4.963 2.038T21 11v1q0 2.925-2.038 4.963T14 19h-4q-.425 0-.713.288T9 20v2H3Z"/>`),
		g.Group(children),
	)
}