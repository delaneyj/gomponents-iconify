package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 76H56a12 12 0 0 1 0-24h136a4 4 0 0 0 0-8H56a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h160a12 12 0 0 0 12-12V88a12 12 0 0 0-12-12Zm4 124a4 4 0 0 1-4 4H56a12 12 0 0 1-12-12V80a19.86 19.86 0 0 0 12 4h160a4 4 0 0 1 4 4Zm-32-60a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}