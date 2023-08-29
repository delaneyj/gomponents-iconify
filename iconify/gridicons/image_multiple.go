package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 7.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 15 7.5zM4 20h14a2 2 0 0 1-2 2H4c-1.1 0-2-.9-2-2V8a2 2 0 0 1 2-2v14zM22 4v12a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2zM8 4v6.333L11 7l4.855 5.395l.656-.731a2 2 0 0 1 2.976 0l.513.57V4H8z"/>`),
		g.Group(children),
	)
}