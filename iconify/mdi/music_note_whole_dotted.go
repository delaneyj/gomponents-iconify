package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteWholeDotted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 15a2 2 0 1 1-2 2a2 2 0 0 1 2-2m0-2a4 4 0 1 0 4 4a4 4 0 0 0-4-4m6.5 4a1.5 1.5 0 1 1-1.5 1.5a1.5 1.5 0 0 1 1.5-1.5Z"/>`),
		g.Group(children),
	)
}