package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamWiredStand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 23q-.425 0-.713-.288T7 22v-2q0-2.1 1.463-3.55T12 15q.275 0 .525.025t.5.075l.55-.85l-1.85-.175q-2.425-.25-4.075-2.063T6 7.75q0-2.475 1.638-4.288T11.724 1.4L15.85 1q.875-.075 1.525.525T18.025 3v9.475q0 .875-.65 1.475t-1.525.525h-.025l-.95 1.425q.975.675 1.55 1.75T17 20v2q0 .425-.288.713T16 23H8Z"/>`),
		g.Group(children),
	)
}