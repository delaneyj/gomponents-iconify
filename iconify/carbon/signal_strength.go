package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalStrength(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 6h2v20h-2zm-8 5h2v15h-2zm-7 5h2v10h-2zm-7 4h2v6H4z"/>`),
		g.Group(children),
	)
}