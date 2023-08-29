package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playlist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 17a3 3 0 1 0 6 0a3 3 0 1 0-6 0m6 0V4h4m-8 1H3m0 4h10m-4 4H3"/>`),
		g.Group(children),
	)
}