package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardAltWindowsKeyboardKeyAltPc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2.5v-1a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v1m0 9v1a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-.75M10.5 5h3M12 5v4M6.5 5v4h2m-5 0V6.5a1.5 1.5 0 0 0-3 0V9m0-1.5h3"/>`),
		g.Group(children),
	)
}