package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M12.148 14.052a.5.5 0 0 1-.543.453l-5.186-.472a.5.5 0 1 1 .09-.995l5.186.471a.5.5 0 0 1 .453.543Z"/><path d="M5.948 7.852a.5.5 0 0 1 .543.453l.471 5.185a.5.5 0 1 1-.995.09l-.472-5.185a.5.5 0 0 1 .453-.543Z"/><path d="M6.646 13.354a.5.5 0 0 1 0-.708l6.536-6.535a.5.5 0 0 1 .707.707l-6.535 6.536a.5.5 0 0 1-.708 0Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}