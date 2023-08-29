package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.753 22.91C10.028 18.639 14.93 5.283 28.775 5.283c6.923 0 9.649 3.051 9.649 8.608c0 5.735-6.479 13.467-19.747 13.467m7.517-17.532L11.421 44.283"/>`),
		g.Group(children),
	)
}