package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeOrangeDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#f4aa41" d="m36 10.984l24.974 24.974L36 60.932L11.026 35.958z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36 10.984l24.974 24.974L36 60.932L11.026 35.958z"/>`),
		g.Group(children),
	)
}