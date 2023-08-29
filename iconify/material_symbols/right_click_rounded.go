package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightClickRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.3 18l.625-2.1q1.35-.325 2.212-1.4T16 12q0-1.65-1.175-2.825T12 8q-1.425 0-2.5.863t-1.4 2.212L6 11.7q.125-2.4 1.85-4.05T12 6q2.5 0 4.25 1.75T18 12q0 2.425-1.65 4.15T12.3 18Zm-4.55.225l-3.3 3.3q-.425.425-.975.413T2.5 21.5q-.425-.425-.425-.988t.425-.987l3.275-3.275L3.5 15.5q-.35-.125-.337-.475t.362-.475l7.575-2.275q.275-.1.5.125t.125.5L9.45 20.475q-.125.35-.475.363T8.5 20.5l-.75-2.275Z"/>`),
		g.Group(children),
	)
}