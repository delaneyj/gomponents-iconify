package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWalletTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M4 8h40v8s-10 2-10 8s10 8 10 8v8H4V8Z"/><path stroke-linecap="round" d="M44 24h-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWalletTwo0)"/>`),
		g.Group(children),
	)
}