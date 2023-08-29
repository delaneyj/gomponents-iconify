package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM292.969 367.896h614.062v86.353L600 629.297L292.969 454.248v-86.352zm0 141.577l156.372 89.136l-156.372 163.915V509.473zm614.062 0v253.052L750.659 598.608l156.372-89.135zM492.334 623.145L600 684.521l107.666-61.377l199.365 208.96H292.969l199.365-208.959z"/>`),
		g.Group(children),
	)
}