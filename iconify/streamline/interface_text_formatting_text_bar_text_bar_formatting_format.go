package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTextBarTextBarFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 13.5H9a2 2 0 0 1-2-2a2 2 0 0 1-2 2H3.5m7-13H9a2 2 0 0 0-2 2a2 2 0 0 0-2-2H3.5m3.5 2v9M4.5 8h5"/>`),
		g.Group(children),
	)
}