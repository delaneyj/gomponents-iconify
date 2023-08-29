package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6a3 3 0 0 1 3-3h16.75A3.25 3.25 0 0 1 26 6.25v1.006c1.748.618 3 2.285 3 4.244v13a4.5 4.5 0 0 1-4.5 4.5h-17A4.5 4.5 0 0 1 3 24.5V6.25h.01A3.04 3.04 0 0 1 3 6Zm21 .25C24 5.56 23.44 5 22.75 5H6a1 1 0 0 0 0 2h18v-.75ZM21 18a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3Z"/>`),
		g.Group(children),
	)
}