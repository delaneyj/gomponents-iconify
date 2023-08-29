package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.354 16.354a.5.5 0 0 1-.708 0l-4-4a.5.5 0 0 1 0-.708l4-4a.5.5 0 1 1 .708.708L3.707 12l3.647 3.646a.5.5 0 0 1 0 .708Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.75 11.5a2.74 2.74 0 0 0 1.943-.81c.516-.52.807-1.226.807-1.963V3a.5.5 0 0 1 1 0v5.727c0 1-.394 1.959-1.097 2.667A3.739 3.739 0 0 1 12.75 12.5H3a.5.5 0 0 1 0-1h9.75Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}