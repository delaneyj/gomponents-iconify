package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IonicSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 2.25a5.75 5.75 0 1 0 5.723 5.19a.75.75 0 0 1 1.493-.144a7.25 7.25 0 1 1-3.48-5.51a.75.75 0 1 1-.773 1.285A5.72 5.72 0 0 0 8 2.25Z"/><path d="M8 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 1.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm5.5-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 1.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/></g>`),
		g.Group(children),
	)
}