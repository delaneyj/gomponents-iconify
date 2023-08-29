package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicArtistThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="7" r="3"/><circle cx="18" cy="18" r="2" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.341 20H6a2 2 0 0 1-2-2a4 4 0 0 1 4-4h5.528M20 18v-7l2 2"/></g>`),
		g.Group(children),
	)
}