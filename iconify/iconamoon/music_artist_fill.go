package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicArtistFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm3 15a3 3 0 0 1 4-2.83V11a1 1 0 0 1 1.707-.707l2 2a1 1 0 0 1-1.414 1.414L21 13.414V18a3 3 0 1 1-6 0ZM3 18a5 5 0 0 1 5-5h5.528a1 1 0 1 1 0 2H8a3 3 0 0 0-3 3a1 1 0 0 0 1 1h6.341a1 1 0 1 1 0 2H6a3 3 0 0 1-3-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}