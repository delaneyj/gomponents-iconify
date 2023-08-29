package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnoozeWeekend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.5 14v-4a3 3 0 0 1 3-3h15a3 3 0 0 1 3 3v4m-21 0A2.5 2.5 0 0 0 3 16.5V25h5M5.5 14A2.5 2.5 0 0 1 8 16.5V20h16v-3.5a2.5 2.5 0 0 1 2.5-2.5m0 0a2.5 2.5 0 0 1 2.5 2.5V25h-5m0 0H8m16 0v2M8 25v2"/>`),
		g.Group(children),
	)
}