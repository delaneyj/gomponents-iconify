package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardComputerKeyboardDeviceElectronicsDvorakQwerty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="5.5" rx="1"/><path d="M3.5 10.75h7m-7-2.5h7m-7-2.75v-1a2 2 0 0 1 2-2h4a2 2 0 0 0 2-2"/></g>`),
		g.Group(children),
	)
}