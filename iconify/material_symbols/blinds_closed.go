package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlindsClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22.75q-.725 0-1.238-.513T13.25 21H2v-2h2V3h16v16h2v2h-5.25q0 .725-.513 1.238T15 22.75ZM6 7h8V5H6v2Zm10 0h2V5h-2v2ZM6 11h8V9H6v2Zm10 0h2V9h-2v2ZM6 15h8v-2H6v2Zm10 0h2v-2h-2v2ZM6 19h8v-2H6v2Zm10 0h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}