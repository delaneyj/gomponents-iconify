package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 3A1.5 1.5 0 0 0 3 4.5v10A2.5 2.5 0 0 0 5.5 17H15a2 2 0 0 0 2-2V7a2 2 0 0 0-1-1.732V5a2 2 0 0 0-2-2H4.5ZM15 5H4.5a.5.5 0 0 1 0-1H14a1 1 0 0 1 1 1Zm-1.5 6h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}