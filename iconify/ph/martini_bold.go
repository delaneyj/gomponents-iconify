package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MartiniBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M243.09 35.41A12 12 0 0 0 232 28H24a12 12 0 0 0-8.48 20.49L116 149v55H88a12 12 0 0 0 0 24h80a12 12 0 0 0 0-24h-28v-55L240.48 48.49a12 12 0 0 0 2.61-13.08ZM203 52l-12 12H65L53 52Zm-75 75L89 88h78Z"/>`),
		g.Group(children),
	)
}