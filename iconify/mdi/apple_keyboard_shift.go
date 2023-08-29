package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppleKeyboardShift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 18v-6h2.17L12 6.83L6.83 12H9v6h6M12 4l10 10h-5v6H7v-6H2L12 4Z"/>`),
		g.Group(children),
	)
}