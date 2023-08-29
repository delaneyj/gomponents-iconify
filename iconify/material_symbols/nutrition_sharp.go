package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutritionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-2.925 0-4.963-2.038T5 14q0-2.35 1.375-4.2t3.65-2.5q-.525-.125-1-.375T8.2 6.3q-.825-.8-1.075-1.937t-.1-2.338Q8.2 1.9 9.338 2.15T11.3 3.2q.55.575.813 1.313t.337 1.512q.425-.975 1.075-1.838t1.5-1.537l1.4 1.4q-.8.65-1.425 1.463t-.9 1.812q2.2.7 3.55 2.538T19 14q0 2.925-2.05 4.963T12 21Z"/>`),
		g.Group(children),
	)
}