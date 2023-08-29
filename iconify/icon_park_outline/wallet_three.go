package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M39 16V9a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h27a3 3 0 0 0 3-3v-7"/><rect width="20" height="16" x="22" y="16" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="3"/><circle r="2" fill="currentColor" transform="matrix(0 -1 -1 0 30 24)"/></g>`),
		g.Group(children),
	)
}