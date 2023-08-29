package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Padland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.75 42.5L5.5 30.366V5.5l9.25 12.134V42.5zm9.25 0l-9.25-12.134V5.5L24 17.634V42.5zm9.25 0L24 30.366V5.5l9.25 12.134V42.5zm9.25 0l-9.25-12.134V5.5l9.25 12.134V42.5z"/>`),
		g.Group(children),
	)
}