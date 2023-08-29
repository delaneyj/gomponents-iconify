package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16q.625 0 1.063-.425T18.5 14.5q0-.625-.438-1.063T17 13q-.65 0-1.075.438T15.5 14.5q0 .65.425 1.075T17 16ZM2 9l3.4-3.4q.275-.275.638-.438T6.825 5H17.15q.425 0 .788.163t.637.437L22 9H2Zm2 10q-.85 0-1.425-.575T2 17v-6h20v6q0 .85-.588 1.425T20 19H4Z"/>`),
		g.Group(children),
	)
}