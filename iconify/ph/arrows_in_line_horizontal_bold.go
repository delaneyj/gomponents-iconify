package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInLineHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 40v176a12 12 0 0 1-24 0V40a12 12 0 0 1 24 0ZM64.49 87.51a12 12 0 0 0-17 17L59 116H16a12 12 0 0 0 0 24h43l-11.49 11.51a12 12 0 0 0 17 17l32-32a12 12 0 0 0 0-17ZM240 116h-43l11.52-11.51a12 12 0 0 0-17-17l-32 32a12 12 0 0 0 0 17l32 32a12 12 0 0 0 17-17L197 140h43a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}