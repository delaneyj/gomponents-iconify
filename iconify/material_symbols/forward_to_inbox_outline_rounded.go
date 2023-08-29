package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardToInboxOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6H4v12h9v2H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v7h-2V6ZM4 6v12v-5v.075V6Zm8 5l8-5v2l-7.475 4.675q-.25.15-.525.15t-.525-.15L4 8V6l8 5Zm7.175 9H16q-.425 0-.713-.288T15 19q0-.425.288-.713T16 18h3.175l-.9-.9q-.3-.3-.287-.7t.312-.7q.3-.275.7-.287t.7.287l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T18.3 22.3q-.275-.275-.275-.7t.275-.7l.875-.9Z"/>`),
		g.Group(children),
	)
}