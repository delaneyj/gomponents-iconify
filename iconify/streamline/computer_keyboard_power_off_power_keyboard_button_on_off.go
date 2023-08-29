package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardPowerOffPowerKeyboardButtonOnOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M5.5 3.5h-2v2m0-2L7 7m1.57-3.13a3.5 3.5 0 1 1-4.7 4.7"/></g>`),
		g.Group(children),
	)
}