package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperBoardOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.825 17L16 13.15V11h-2.15l-1.025-1H16V7h-4v2.175l-1-1V7H9.85L5.825 3H20v4h2v2h-2v2h2v2h-2v2h2v2h-2.175ZM6 17h5v-4H6v4ZM3.2 3.175l16.8 16.8V21H2V3.175h1.2Zm8.8 8.8V17h4v-1.025l-4-4ZM7.025 7H6v5h5v-1.025L7.025 7Zm13.45 16.3L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425Z"/>`),
		g.Group(children),
	)
}