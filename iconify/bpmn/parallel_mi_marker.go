package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParallelMiMarker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M500 400v1200h200V400H500zm400 0v1200h200V400H900zm400 0v1200h200V400h-200z"/>`),
		g.Group(children),
	)
}