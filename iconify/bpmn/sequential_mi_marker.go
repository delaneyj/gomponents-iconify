package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SequentialMiMarker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M400 500v200h1200V500H400zm0 400v200h1200V900H400zm0 400v200h1200v-200H400z"/>`),
		g.Group(children),
	)
}