package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 24C4 12.954 12.954 4 24 4s20 8.954 20 20s-8.954 20-20 20S4 35.046 4 24Zm22-11a2 2 0 1 0-4 0v14a2 2 0 1 0 4 0V13Zm-2 20a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}