package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWalletOne0"><g fill="none"><rect width="38" height="26" x="5" y="13" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M25 20h18v11H25z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M43 18v15M32 13c0-5-3.5-6-5-6H12.5C8.9 7 8 9.866 8 11.299V13"/><circle cx="33.5" cy="25.5" r="1.5" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWalletOne0)"/>`),
		g.Group(children),
	)
}