package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 18h-6a2.002 2.002 0 0 0-2 2v2h-5.586L10 17.586V12h2a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2H6a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h2v5.586l-4.707 4.707a1 1 0 0 0 0 1.414l5 5a1 1 0 0 0 1.414 0L14.414 24H20v2a2.002 2.002 0 0 0 2 2h6a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2ZM6 4h6v6H6Zm3 22.586L5.414 23L9 19.414L12.586 23ZM22 26v-6h6v6Z"/>`),
		g.Group(children),
	)
}