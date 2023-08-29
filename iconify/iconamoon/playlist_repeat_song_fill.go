package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistRepeatSongFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 9a1 1 0 1 1-2 0V5.618l-.553.276a1 1 0 1 1-.894-1.788l2-1A1 1 0 0 1 22 4v5ZM6 19a1 1 0 0 1-1.707.707l-2-2a1 1 0 0 1 0-1.414l2-2A1 1 0 0 1 6 15v1h11a3 3 0 0 0 3-3a1 1 0 1 1 2 0a5 5 0 0 1-4.998 5H6v1ZM16 7a1 1 0 0 0-1-1H7a5 5 0 0 0-5 5a1 1 0 1 0 2 0a3 3 0 0 1 3-3h8a1 1 0 0 0 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}