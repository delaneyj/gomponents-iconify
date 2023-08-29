package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfBboxPolygon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M85.445 87.445h-70.89a2 2 0 0 1-2-2v-70.89a2 2 0 0 1 2-2h70.891a2 2 0 0 1 2 2v70.891a2 2 0 0 1-2.001 1.999zm-68.89-4h66.891v-66.89H16.555v66.89z"/>`),
		g.Group(children),
	)
}