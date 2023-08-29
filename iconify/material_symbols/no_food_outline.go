package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoFoodOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.625 18.75L19.8 16.925L20.8 7h-9.55L11 5h5V1h2v4h5l-1.375 13.75Zm-6-5.975Zm4.85 10.525L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM1 19v-2h15v2H1Zm1 4q-.425 0-.713-.288T1 22v-1h15v1q0 .425-.288.713T15 23H2ZM9.05 9.025v2q-.125 0-.275-.013T8.5 11q-1.475 0-2.788.5T3.676 13h9.35l2 2H1q0-3.025 2.337-4.513T8.5 9q.125 0 .275.013t.275.012ZM8.5 13Z"/>`),
		g.Group(children),
	)
}