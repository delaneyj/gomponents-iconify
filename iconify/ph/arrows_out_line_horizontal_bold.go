package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutLineHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 40v176a12 12 0 0 1-24 0V40a12 12 0 0 1 24 0Zm-52 76H45l11.52-11.51a12 12 0 0 0-17-17l-32 32a12 12 0 0 0 0 17l32 32a12 12 0 0 0 17-17L45 140h43a12 12 0 0 0 0-24Zm160.49 3.51l-32-32a12 12 0 0 0-17 17L211 116h-43a12 12 0 0 0 0 24h43l-11.52 11.51a12 12 0 0 0 17 17l32-32a12 12 0 0 0 .01-17Z"/>`),
		g.Group(children),
	)
}