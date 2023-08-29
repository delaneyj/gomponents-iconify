package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M15.148 17.052a.5.5 0 0 1-.543.453l-5.186-.471a.5.5 0 1 1 .09-.996l5.186.471a.5.5 0 0 1 .453.543Z"/><path d="M8.948 10.852a.5.5 0 0 1 .543.453l.471 5.185a.5.5 0 1 1-.995.09l-.472-5.185a.5.5 0 0 1 .453-.543Z"/><path d="M9.646 16.354a.5.5 0 0 1 0-.708l6.536-6.535a.5.5 0 0 1 .707.707l-6.535 6.536a.5.5 0 0 1-.708 0Z"/></g><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}