package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.005 10v10a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1V10h20Zm0-2h-20V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v4Zm-7 8v2h4v-2h-4Z"/>`),
		g.Group(children),
	)
}