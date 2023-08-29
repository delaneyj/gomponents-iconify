package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M11.5 17.999a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm-1-10.5h-3a1.501 1.501 0 0 1 3-.001Z"/><path d="M7.5 7.999a.5.5 0 1 1 0-1h11a.5.5 0 0 1 0 1h-11Z"/><path fill-rule="evenodd" d="M17.5 8.5h-9A.5.5 0 0 0 8 9v11a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V9a.5.5 0 0 0-.5-.5ZM9 19.5v-10h8v10H9Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}