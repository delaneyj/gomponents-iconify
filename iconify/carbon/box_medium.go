package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 28H6a2.002 2.002 0 0 1-2-2V9h2v17h20V9h2v17a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="m18 9l-1.515 5L16 15.977L15.535 14L14 9h-2v14h2v-8l-.158-1.996l.579 1.996L16 19.626L17.579 15l.58-2L18 15v8h2V9h-2zM4 4h24v2H4z"/>`),
		g.Group(children),
	)
}