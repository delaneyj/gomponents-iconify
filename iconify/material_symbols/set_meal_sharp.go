package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetMealSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 15V2h22v13H1Zm2.075 3.5L3 17l17.975-.95l.075 1.5l-17.975.95ZM3 20.975v-1.5h18v1.5H3Zm7.25-9.475q1.95 0 3.638-.675T16.8 8.85q.15 1.025 1.1 1.588T20 11V6q-1.175 0-2.113.588T16.8 8.2q-1.35-1.35-2.988-2.025T10.25 5.5q-2.15 0-3.975.813T3.5 8.5q.95 1.375 2.775 2.188t3.975.812Z"/>`),
		g.Group(children),
	)
}