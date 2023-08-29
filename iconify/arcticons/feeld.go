package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feeld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="14" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7"/><rect width="14" height="37" x="17" y="-6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7" transform="rotate(-90 24 12.5)"/><rect width="14" height="37" x="13.632" y="2.172" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7" transform="rotate(-45 20.632 20.672)"/>`),
		g.Group(children),
	)
}