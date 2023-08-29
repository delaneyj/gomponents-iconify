package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntiClockwiseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.145.146l.708.708l-1.149 1.148A6.499 6.499 0 0 1 14 8.495a6.499 6.499 0 0 1-6.5 6.496A6.499 6.499 0 0 1 1 8.495v-.5h1v.5a5.499 5.499 0 0 0 5.5 5.496A5.5 5.5 0 0 0 13 8.495a5.499 5.499 0 0 0-5.289-5.492l1.142 1.14l-.708.708l-2.352-2.352L8.145.146Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}