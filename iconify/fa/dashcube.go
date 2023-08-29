package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashcube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M0 680q0-131 91.5-226.5T314 358h742L1408 0v1470q0 132-91.5 227t-222.5 95H314q-131 0-222.5-95T0 1470V680zm1232 754l-176-180V829q0-46-32-79t-78-33H462q-46 0-78 33t-32 79v492q0 46 32.5 79.5T462 1434h770z"/>`),
		g.Group(children),
	)
}