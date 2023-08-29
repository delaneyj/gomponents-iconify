package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCardTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.005 11v9a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-9h20Zm0-4h-20V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v3Z"/>`),
		g.Group(children),
	)
}