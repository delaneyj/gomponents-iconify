package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardCommandKeyboardMacCommandApple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.75 10A1.75 1.75 0 1 1 10 11.75v-9.5A1.75 1.75 0 1 1 11.75 4h-9.5A1.75 1.75 0 1 1 4 2.25v9.5A1.75 1.75 0 1 1 2.25 10Z"/>`),
		g.Group(children),
	)
}