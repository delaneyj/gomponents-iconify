package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M23.8 14C29.433 14 34 18.477 34 24s-4.567 10-10.2 10c-2.612 0-4.995-.963-6.8-2.546"/><path d="M24 4c11.046 0 20 8.954 20 20s-8.954 20-20 20a19.934 19.934 0 0 1-13.927-5.646A19.94 19.94 0 0 1 4 24a19.933 19.933 0 0 1 5.556-13.833"/><path d="M9.556 10.167L24 24L10.073 38.354"/></g>`),
		g.Group(children),
	)
}