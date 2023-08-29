package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsHockey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-3q0-.425.288-.713T3 16h1v4H2Zm3 0v-4h4l.85-1.95l1.6 3.5l-.85 1.9q-.125.275-.375.413T9.7 20H5Zm15 0v-4h1q.425 0 .713.288T22 17v3h-2Zm-1 0h-4.7q-.275 0-.525-.138t-.375-.412L6.35 4H9.7L12 9.2L14.3 4h3.35l-4.05 8.85L15 16h4v4Z"/>`),
		g.Group(children),
	)
}