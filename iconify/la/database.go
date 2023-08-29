package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 4v24h20V4zm2 2h16v5H8zm0 7h16v6H8zm0 8h16v5H8z"/>`),
		g.Group(children),
	)
}