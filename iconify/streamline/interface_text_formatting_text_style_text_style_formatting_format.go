package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTextStyleTextStyleFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5v13m-2.5 0h5M1.5 3V1.5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1V3"/>`),
		g.Group(children),
	)
}