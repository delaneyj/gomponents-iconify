package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApertureFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 128A104 104 0 0 0 54.46 54.46A104 104 0 0 0 128 232h.09A104 104 0 0 0 232 128ZM49.18 88.92l51.21 9.35l-53.74 63.26a88.39 88.39 0 0 1 2.53-72.61Zm160.17 5.54a88.41 88.41 0 0 1-2.53 72.62l-51.21-9.35Zm-8.08-15.2L167.55 119l-27.92-78.22a87.38 87.38 0 0 1 50.6 25a88.74 88.74 0 0 1 11.04 13.48Zm-78.84-39.07l17.51 49L58.3 74.32a89.28 89.28 0 0 1 7.47-8.55a87.37 87.37 0 0 1 56.66-25.58Zm-67.7 136.55L88.45 137l27.92 78.18a88 88 0 0 1-61.64-38.48Zm78.84 39.06l-17.51-49l23.08 4.2l58.52 10.69a87.5 87.5 0 0 1-64.13 34.12Z"/>`),
		g.Group(children),
	)
}