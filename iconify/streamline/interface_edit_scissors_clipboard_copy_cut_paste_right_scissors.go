package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditScissorsClipboardCopyCutPasteRightScissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.19 4.93l11.31 6.52"/><circle cx="2.75" cy="2.75" r="2.25"/><path d="M2.19 9.07L13.5 2.55"/><circle cx="2.75" cy="11.25" r="2.25"/></g>`),
		g.Group(children),
	)
}