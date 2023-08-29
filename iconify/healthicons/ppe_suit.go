package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 8a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v6.417h1a4 4 0 0 1 4 4V31a3 3 0 0 1-4 2.83V41a3 3 0 1 1-6 0V29h-2v12a3 3 0 1 1-6 0v-7.17A3 3 0 0 1 13 31V18.417a4 4 0 0 1 4-4h1V8Zm13 23a1 1 0 1 0 2 0v-1h-2v1Zm-15 1a1 1 0 0 1-1-1v-1h2v1a1 1 0 0 1-1 1Zm10-21a2 2 0 0 0 2-2V7h-8v2a2 2 0 0 0 2 2h4Zm-8.6 12.792h5.5v-9.375h2.2v9.375h5.5v2.083H17.4v-2.083Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}