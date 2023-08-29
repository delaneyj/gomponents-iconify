package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionSharpTurnFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-2 24H16v-2h6.586L8 9.414V26H6V7a1 1 0 0 1 1.707-.707L24 22.586V16h2Z"/><path fill="none" d="M26 26H16v-2h6.586L8 9.414V26H6V7a1 1 0 0 1 1.707-.707L24 22.586V16h2Z"/>`),
		g.Group(children),
	)
}