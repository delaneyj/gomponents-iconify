package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M8.5 14.999a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm-1-10.5h-3a1.501 1.501 0 0 1 3-.001Z"/><path d="M4.5 4.999a.5.5 0 1 1 0-1h11a.5.5 0 0 1 0 1h-11Z"/><path fill-rule="evenodd" d="M14.5 5.5h-9A.5.5 0 0 0 5 6v11a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5ZM6 16.5v-10h8v10H6Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}