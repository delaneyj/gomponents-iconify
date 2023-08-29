package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tactic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.4 21L15 19.6l1.575-1.6L15 16.425L16.4 15l1.6 1.6l1.575-1.6L21 16.425L19.4 18l1.6 1.6l-1.425 1.4L18 19.425L16.4 21ZM6 19q.425 0 .713-.288T7 18q0-.425-.288-.713T6 17q-.425 0-.713.288T5 18q0 .425.288.713T6 19Zm0 2q-1.25 0-2.125-.875T3 18q0-1.25.875-2.125T6 15q.925 0 1.688.513T8.8 16.9q.975-.275 1.588-1.075T11 14v-4q0-2.075 1.463-3.538T16 5h1.15l-1.575-1.575L17 2l4 4l-4 4l-1.425-1.4L17.15 7H16q-1.25 0-2.125.875T13 10v4q0 1.825-1.175 3.213T8.85 18.924q-.3.925-1.088 1.5T6 21ZM4.4 9L3 7.6L4.575 6L3 4.425L4.4 3L6 4.6L7.575 3L9 4.425L7.4 6L9 7.6L7.575 9L6 7.425L4.4 9Z"/>`),
		g.Group(children),
	)
}