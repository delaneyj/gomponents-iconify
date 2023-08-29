package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoginOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 3.25a.75.75 0 0 0 0 1.5a7.25 7.25 0 1 1 0 14.5a.75.75 0 0 0 0 1.5a8.75 8.75 0 1 0 0-17.5Z"/><path d="M10.47 9.53a.75.75 0 0 1 1.06-1.06l3 3a.75.75 0 0 1 0 1.06l-3 3a.75.75 0 1 1-1.06-1.06l1.72-1.72H4a.75.75 0 0 1 0-1.5h8.19l-1.72-1.72Z"/></g>`),
		g.Group(children),
	)
}