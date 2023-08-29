package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 3.5v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1H3m-2.5 0v-1a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v1H3m-2.5 0H3m6 6h3"/>`),
		g.Group(children),
	)
}