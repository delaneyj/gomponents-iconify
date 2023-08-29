package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func V(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 38a2 2 0 0 0 1.846-1.23l10-24a2 2 0 0 0-3.692-1.54L24 30.8l-8.154-19.57a2 2 0 0 0-3.692 1.54l10 24A2 2 0 0 0 24 38Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}