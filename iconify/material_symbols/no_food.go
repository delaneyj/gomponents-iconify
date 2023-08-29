package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.625 18.75L11.45 8.6L11 5h5V1h2v4h5l-1.375 13.75Zm-1.15 4.55l-8.3-8.3H1q0-3.025 2.337-4.513T8.5 9q.125 0 .263.013t.287.012l-.025 2.825L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM1 19v-2h15v2H1Zm1 4q-.425 0-.713-.288T1 22v-1h15v1q0 .425-.288.713T15 23H2Z"/>`),
		g.Group(children),
	)
}