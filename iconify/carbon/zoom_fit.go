package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomFit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.448 20A10.856 10.856 0 0 0 24 13a11 11 0 1 0-11 11a10.856 10.856 0 0 0 7-2.552L27.586 29L29 27.586ZM13 22a9 9 0 1 1 9-9a9.01 9.01 0 0 1-9 9Z"/><path fill="currentColor" d="M10 12H8v-2a2.002 2.002 0 0 1 2-2h2v2h-2zm8 0h-2v-2h-2V8h2a2.002 2.002 0 0 1 2 2zm-6 6h-2a2.002 2.002 0 0 1-2-2v-2h2v2h2zm4 0h-2v-2h2v-2h2v2a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}