package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2H7Z"/><path stroke-linejoin="round" d="M13 3v6h6"/><circle cx="11" cy="17" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12v5m2-3l-2-2"/></g>`),
		g.Group(children),
	)
}