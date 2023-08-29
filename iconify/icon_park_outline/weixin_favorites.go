package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinFavorites(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 14l18-8l18 8M6 14l18 8M6 14v20l18 8m18-28l-18 8m18-8v20l-18 8m0-20v20"/>`),
		g.Group(children),
	)
}