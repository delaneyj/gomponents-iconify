package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListSquareThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.5A4.5 4.5 0 0 1 7.5 3h17A4.5 4.5 0 0 1 29 7.5v17a4.5 4.5 0 0 1-4.5 4.5h-17A4.5 4.5 0 0 1 3 24.5v-17Zm9 3a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm0 5.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm-1.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM14 10.5a1 1 0 0 0 1 1h7a1 1 0 1 0 0-2h-7a1 1 0 0 0-1 1Zm1 4.5a1 1 0 1 0 0 2h7a1 1 0 1 0 0-2h-7Zm-1 6.5a1 1 0 0 0 1 1h7a1 1 0 1 0 0-2h-7a1 1 0 0 0-1 1Z"/>`),
		g.Group(children),
	)
}