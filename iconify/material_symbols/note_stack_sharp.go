package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteStackSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 22l.025-15.025L22 7v10l-5 5H7Zm13-6h-4v4l4-4ZM4.3 19.075L1.675 4.3L16.45 1.675L17.05 5H5v13.95l-.7.125Z"/>`),
		g.Group(children),
	)
}