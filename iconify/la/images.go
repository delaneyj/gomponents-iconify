package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Images(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 5v18h4v4h24V9h-4V5zm2 2h20v14H4zm2 2v10h16V9zm2 2h12v6H8zm18 0h2v14H8v-2h18z"/>`),
		g.Group(children),
	)
}