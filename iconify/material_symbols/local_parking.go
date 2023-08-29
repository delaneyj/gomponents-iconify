package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalParking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21V3h7q2.5 0 4.25 1.75T19 9q0 2.5-1.75 4.25T13 15h-3v6H6Zm4-10h3.2q.825 0 1.413-.588T15.2 9q0-.825-.587-1.413T13.2 7H10v4Z"/>`),
		g.Group(children),
	)
}