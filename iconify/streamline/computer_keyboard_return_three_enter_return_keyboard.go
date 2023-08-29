package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardReturnThreeEnterReturnKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="m5.5 10.5l-2-2l2-2"/><path d="M3.5 8.5h5a1 1 0 0 0 1-1v-3"/></g>`),
		g.Group(children),
	)
}