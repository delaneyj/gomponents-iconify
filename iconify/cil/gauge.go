package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gauge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M425.706 142.294A240 240 0 0 0 16 312v88h144v-32H48v-56c0-114.691 93.309-208 208-208s208 93.309 208 208v56H352v32h144v-88a238.432 238.432 0 0 0-70.294-169.706Z"/><path fill="currentColor" d="M80 264h32v32H80zm160-136h32v32h-32zm-104 40h32v32h-32zm264 96h32v32h-32zm-102.778 71.1l69.2-144.173l-28.85-13.848l-69.183 144.135a64.141 64.141 0 1 0 28.833 13.886ZM256 416a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}