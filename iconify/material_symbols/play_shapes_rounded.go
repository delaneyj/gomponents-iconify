package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayShapesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.7 11q-.575 0-.863-.488t-.012-.987l3.3-5.95q.275-.5.875-.5t.875.5l3.3 5.95q.275.5-.013.988T10.3 11H3.7ZM7 21q-1.65 0-2.825-1.175T3 17q0-1.65 1.175-2.825T7 13q1.65 0 2.825 1.175T11 17q0 1.65-1.175 2.825T7 21Zm7 0q-.425 0-.713-.288T13 20v-6q0-.425.288-.713T14 13h6q.425 0 .713.288T21 14v6q0 .425-.288.713T20 21h-6Zm1.6-14.5l-1.9-1.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275L17 5.1l1.9-1.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-1.9 1.9l1.9 1.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L17 7.9l-1.9 1.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l1.9-1.9Z"/>`),
		g.Group(children),
	)
}