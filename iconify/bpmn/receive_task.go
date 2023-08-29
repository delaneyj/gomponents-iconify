package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiveTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g stroke="currentColor" transform="translate(0 947.638)"><rect width="1800" height="1460" x="100" y="-677.638" fill="transparent" stroke-linecap="round" stroke-width="120" rx="329.651" ry="328.5"/><g fill="none" stroke-width="45.443"><path d="M376.496-429.067h1007.533v649.183H376.496z"/><path stroke-miterlimit="1.4" d="m386.883-429.067l493.38 324.592l493.379-324.592z"/></g></g>`),
		g.Group(children),
	)
}