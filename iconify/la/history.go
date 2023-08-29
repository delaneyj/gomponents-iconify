package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4A11.989 11.989 0 0 0 6 9.344V6H4v7h7v-2H7.375C9.102 8.02 12.297 6 16 6c5.535 0 10 4.465 10 10s-4.465 10-10 10S6 21.535 6 16H4c0 6.617 5.383 12 12 12s12-5.383 12-12S22.617 4 16 4zm-1 4v9h7v-2h-5V8z"/>`),
		g.Group(children),
	)
}