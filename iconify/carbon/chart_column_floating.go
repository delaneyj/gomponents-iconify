package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartColumnFloating(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 22h-8V4h8zm-6-2h4V6h-4zm-6 4H8V10h8zm-6-2h4V12h-4z"/><path fill="currentColor" d="M30 30H4a2.002 2.002 0 0 1-2-2V2h2v26h26Z"/>`),
		g.Group(children),
	)
}