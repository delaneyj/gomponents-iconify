package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v24h2V19h8v3h12V8H17V5zm2 2h8v10H7zm10 3h8v10h-8z"/>`),
		g.Group(children),
	)
}