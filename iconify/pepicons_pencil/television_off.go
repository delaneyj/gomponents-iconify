package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.25 9v5a3 3 0 0 0 3 3h7.5a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-7.5a3 3 0 0 0-3 3Zm3 7a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h7.5a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-7.5Z" clip-rule="evenodd"/><path d="m9.86 5.688l2-2.5c.416-.52 1.197.104.78.624l-2 2.5c-.416.52-1.197-.104-.78-.624Z"/><path d="m8.86 6.312l-2-2.5c-.417-.52.364-1.145.78-.624l2 2.5c.417.52-.364 1.145-.78.624ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}