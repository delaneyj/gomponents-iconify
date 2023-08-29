package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingStrikeThroughTextStrikeThroughFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.08 8.67a2.73 2.73 0 0 1-.42 3.95a5.06 5.06 0 0 1-5.37.21l-.9-.51m6.61-11C8 .33 4.23-.26 3.38 2.54a2.47 2.47 0 0 0 .55 2.38a11.8 11.8 0 0 0 2.6 1.54M.5 6.5h13"/>`),
		g.Group(children),
	)
}