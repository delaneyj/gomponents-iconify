package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWalletThree0"><g fill="none"><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M39 16V9a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h27a3 3 0 0 0 3-3v-7"/><rect width="20" height="16" x="22" y="16" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="3"/><circle r="2" fill="#000" transform="matrix(0 -1 -1 0 30 24)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWalletThree0)"/>`),
		g.Group(children),
	)
}