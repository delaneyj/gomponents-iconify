package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M7.5 7.25a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path fill-rule="evenodd" d="M8.5 7.25a1 1 0 0 1 1-1h7a1 1 0 1 1 0 2h-7a1 1 0 0 1-1-1Zm0 4a1 1 0 0 1 1-1h7a1 1 0 1 1 0 2h-7a1 1 0 0 1-1-1Zm0 4a1 1 0 0 1 1-1h7a1 1 0 1 1 0 2h-7a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g><path d="M6.5 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path fill-rule="evenodd" d="M7.5 6.5A.5.5 0 0 1 8 6h8a.5.5 0 0 1 0 1H8a.5.5 0 0 1-.5-.5Zm0 4A.5.5 0 0 1 8 10h8a.5.5 0 0 1 0 1H8a.5.5 0 0 1-.5-.5Zm0 4A.5.5 0 0 1 8 14h8a.5.5 0 0 1 0 1H8a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}