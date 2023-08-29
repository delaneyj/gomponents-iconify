package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M69.328 17.517H30.567v.01a2.495 2.495 0 0 0-2.396 2.49v59.967a2.495 2.495 0 0 0 2.396 2.489v.011h38.761a2.5 2.5 0 0 0 2.5-2.5V20.017a2.5 2.5 0 0 0-2.5-2.5zM50.059 79.9a2.45 2.45 0 1 1 0-4.9a2.45 2.45 0 0 1 0 4.9zm11.813-7.395H38.128V27.473h23.743v45.032z"/>`),
		g.Group(children),
	)
}