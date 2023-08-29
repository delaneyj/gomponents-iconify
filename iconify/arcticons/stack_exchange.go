package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="30.016" height="32.016" x="8.992" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.041"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.992 20.508h30.016M8.992 28.512h30.016M8.992 12.504h30.016M26.897 36.516V43.5l8.004-6.984"/>`),
		g.Group(children),
	)
}