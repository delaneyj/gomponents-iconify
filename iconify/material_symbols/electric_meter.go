package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricMeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21.95v-2.5q-2.65-.925-4.325-3.237T3 10.95q0-1.875.713-3.512t1.924-2.85q1.213-1.213 2.85-1.925t3.488-.713q1.85 0 3.5.713t2.875 1.925q1.225 1.212 1.938 2.85T21 10.95q0 2.95-1.688 5.238T15 19.425v2.525h-2V19.9q-.25.05-.5.05h-.525q-.25 0-.488-.012T11 19.9v2.05H9ZM8 9h8V7H8v2Zm3.25 8l3-3L13 12.75l1.25-1.25l-1.5-1.5l-3 3L11 14.25L9.75 15.5l1.5 1.5Z"/>`),
		g.Group(children),
	)
}