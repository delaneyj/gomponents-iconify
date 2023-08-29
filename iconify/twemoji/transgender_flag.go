package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#5BCEFA" d="M0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-1.3H0V27z"/><path fill="#F5A9B8" d="M.026 20.5L0 25.8h36v-5.3z"/><path fill="#EEE" d="M0 15.3h36v5.3H0z"/><path fill="#F5A9B8" d="M0 9.902h36V15.4H0z"/><path fill="#5BCEFA" d="M36 9a4 4 0 0 0-4-4H4a4 4 0 0 0-4 4v1.2h36V9z"/>`),
		g.Group(children),
	)
}