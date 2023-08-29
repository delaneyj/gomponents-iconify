package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Auction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 38.483a2.446 2.446 0 0 1-2.446 2.446h0a2.446 2.446 0 1 1 2.446-2.446h0ZM9.491 30.409h20.657M6.536 7.07H23.44l10.42 33.743M15.922 7.071L5.5 40.814"/>`),
		g.Group(children),
	)
}