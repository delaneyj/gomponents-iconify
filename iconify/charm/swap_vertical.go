package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.75 5.75l-3-3l-3 3m3 7.5V2.75m9.5 7.5l-3 3l-3-3m3-7.5v10.5"/>`),
		g.Group(children),
	)
}