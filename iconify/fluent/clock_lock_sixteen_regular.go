package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockLockSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 1 0 12v-1a5 5 0 1 0-4.585-6.999a2.986 2.986 0 0 0-1.162.27A6.003 6.003 0 0 1 8 2Zm-.499 3a.5.5 0 0 1 .492.41l.008.09V8H9.5a.5.5 0 0 1 .09.992L9.5 9H7.501a.5.5 0 0 1-.491-.41L7 8.5v-3a.5.5 0 0 1 .5-.5ZM1.5 9v1H1a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-.5V9a2 2 0 1 0-4 0Zm1 1V9a1 1 0 0 1 2 0v1h-2Zm1 2.25a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}