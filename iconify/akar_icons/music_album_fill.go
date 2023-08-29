package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAlbumFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 5a4 4 0 0 1 4.002-4h13.996A4 4 0 0 1 23 5v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4V5Zm19 7a8 8 0 1 1-16 0a8 8 0 0 1 16 0Zm-8 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}