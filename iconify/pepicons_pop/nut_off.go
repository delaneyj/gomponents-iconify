package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15.156 16.992A2 2 0 0 1 13.42 18H6.58a2 2 0 0 1-1.736-1.008l-3.429-6a2 2 0 0 1 0-1.984l3.429-6A2 2 0 0 1 6.58 2h6.84a2 2 0 0 1 1.736 1.008l3.429 6a2 2 0 0 1 0 1.984l-3.429 6ZM6.58 16h6.84l3.428-6l-3.428-6H6.58l-3.428 6l3.428 6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}