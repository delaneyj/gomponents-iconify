package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAlbumThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="2"/><circle cx="18" cy="9" r="2"/><path d="M15.318 3.631a9 9 0 1 0 5.368 10.736M20 9V2l2 2"/></g>`),
		g.Group(children),
	)
}