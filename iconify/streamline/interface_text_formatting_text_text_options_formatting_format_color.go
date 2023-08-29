package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTextTextOptionsFormattingFormatColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 13.5h13M2.14 11L6.5.83a.54.54 0 0 1 1 0L11.86 11M4.07 6.5h5.86"/>`),
		g.Group(children),
	)
}