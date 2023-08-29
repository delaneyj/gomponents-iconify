package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoRightRoundBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16.47 3.47a.75.75 0 0 1 1.06 0l3 3a.75.75 0 0 1 0 1.06l-3 3a.75.75 0 1 1-1.06-1.06L18.94 7l-2.47-2.47a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/><path d="M18.19 6.25H9a5.75 5.75 0 0 0 0 11.5h7a.75.75 0 0 0 0-1.5H9a4.25 4.25 0 0 1 0-8.5h9.19l.75-.75l-.75-.75Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}