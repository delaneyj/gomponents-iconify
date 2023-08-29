package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowModeler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 10H2V2h8zM4 8h4V4H4zm26 22h-8v-8h8zm-6-2h4v-4h-4zm-4-1H8a6 6 0 0 1 0-12v2a4 4 0 0 0 0 8h12zm4-10v-2a4 4 0 0 0 0-8H12V5h12a6 6 0 0 1 0 12z"/><path fill="currentColor" d="M19 11h-6l-3 4l6 6l6-6Z"/>`),
		g.Group(children),
	)
}