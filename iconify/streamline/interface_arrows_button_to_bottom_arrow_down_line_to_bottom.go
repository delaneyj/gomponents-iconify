package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsButtonToBottomArrowDownLineToBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 5.71l6.15 6.14a.48.48 0 0 0 .7 0l6.15-6.14M.5 2h13"/>`),
		g.Group(children),
	)
}