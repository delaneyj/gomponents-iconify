package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audiolab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.313 35.45a8.07 8.07 0 1 0 0 .116"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.312 35.566L26.316 4.5l11.505.002v7.636H26.316"/>`),
		g.Group(children),
	)
}