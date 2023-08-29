package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Product(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 18h6v2H8zm0 4h10v2H8z"/><path fill="currentColor" d="M26 4H6a2.002 2.002 0 0 0-2 2v20a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2Zm-8 2v4h-4V6ZM6 26V6h6v6h8V6h6l.001 20Z"/>`),
		g.Group(children),
	)
}