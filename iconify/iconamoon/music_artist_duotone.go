package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicArtistDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="7" r="3" fill="currentColor" opacity=".16"/><circle cx="12" cy="7" r="3" stroke="currentColor" stroke-width="2"/><circle cx="18" cy="18" r="2" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.341 20H6a2 2 0 0 1-2-2a4 4 0 0 1 4-4h5.528M20 18v-7l2 2"/></g>`),
		g.Group(children),
	)
}