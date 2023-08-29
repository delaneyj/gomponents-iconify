package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataFunnelTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.248 7.504a2.752 2.752 0 0 0 0-5.504H4.75a2.752 2.752 0 1 0 0 5.504h14.498Zm-2 7a2.752 2.752 0 1 0 0-5.504H6.75a2.752 2.752 0 0 0 0 5.504h10.498ZM17 18.752A2.752 2.752 0 0 0 14.248 16H9.75a2.752 2.752 0 0 0 0 5.504h4.498A2.752 2.752 0 0 0 17 18.752Z"/>`),
		g.Group(children),
	)
}