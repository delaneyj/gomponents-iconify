package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppleKeyboardCaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 14V8h2.17L12 2.83L6.83 8H9v6h6M12 0l10 10h-5v6H7v-6H2L12 0M7 18h10v6H7v-6m8 2H9v2h6v-2Z"/>`),
		g.Group(children),
	)
}