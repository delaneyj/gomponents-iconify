package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M1.5 2a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-13a2 2 0 0 1-2-2V2Zm15 0h-13v16h13V2Z"/><path d="M7.75 16a1 1 0 0 1 1-1h2.5a1 1 0 1 1 0 2h-2.5a1 1 0 0 1-1-1Z"/></g>`),
		g.Group(children),
	)
}