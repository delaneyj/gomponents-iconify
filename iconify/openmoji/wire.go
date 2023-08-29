package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#E27022" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M16 36h-5m50 0h-5"/><path fill="#D22F27" d="M16 35.94h40V40H16z"/><path fill="#EA5A47" d="M16 32h40v4H16z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M16 32h40v8H16zm0 4h-5m50 0h-5"/><path fill="#E27022" d="M15 37.05h-4a1.05 1.05 0 1 1 0-2.1h4m42 0h4a1.05 1.05 0 1 1 0 2.1h-4"/>`),
		g.Group(children),
	)
}