package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.82 3.5v41m-7.23-32.517v24.034M9.374 19.759v8.482m21.725-16.258v24.034m7.527-16.258v8.482"/>`),
		g.Group(children),
	)
}