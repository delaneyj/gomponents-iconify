package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignLayersOneDesignLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.47 6.9a1.18 1.18 0 0 1-.94 0L.83 4.26a.56.56 0 0 1 0-1L6.53.6a1.18 1.18 0 0 1 .94 0l5.7 2.64a.56.56 0 0 1 0 1Zm6.03.45l-6.1 2.81a1 1 0 0 1-.83 0L.5 7.35m13 3.25l-6.1 2.81a1 1 0 0 1-.83 0L.5 10.6"/>`),
		g.Group(children),
	)
}