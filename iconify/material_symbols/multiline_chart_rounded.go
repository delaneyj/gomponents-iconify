package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultilineChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.475 12.35L4.1 17.725q-.3.3-.712.288t-.713-.313q-.275-.3-.287-.7t.287-.7l6.1-6.1q.3-.3.7-.3t.7.3l3.3 3.3l2.9-3.25q-1.275-1.5-3-2.375T9.575 7q-1.5 0-2.875.475t-2.525 1.35q-.375.275-.813.263t-.687-.363q-.275-.375-.187-.813T2.95 7.2q1.425-1.05 3.1-1.625T9.575 5q2.45 0 4.525.988t3.625 2.762l2.125-2.425q.3-.35.738-.338t.737.288q.275.275.288.675t-.263.7l-2.375 2.7q.7 1.1 1.163 2.35t.687 2.65q.075.45-.175.8t-.75.35q-.475 0-.725-.25t-.375-.775q-.2-.95-.513-1.837t-.762-1.688l-3.25 3.65q-.275.325-.713.338t-.737-.288l-3.35-3.3Z"/>`),
		g.Group(children),
	)
}