package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAlbumFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.707 1.293A1 1 0 0 0 19 2v4.17A3 3 0 1 0 21 9V4.414l.293.293a1 1 0 1 0 1.414-1.414l-2-2ZM12 4a8 8 0 1 0 7.72 10.105a1 1 0 1 1 1.93.524C20.497 18.876 16.614 22 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2c1.3 0 2.545.249 3.687.702a1 1 0 0 1-.738 1.859A7.976 7.976 0 0 0 12 4Zm0 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}