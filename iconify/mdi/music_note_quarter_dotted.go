package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteQuarterDotted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13v.56a3.96 3.96 0 0 0-2-.56a4 4 0 1 0 4 4V3h-2m4.5 14a1.5 1.5 0 1 1-1.5 1.5a1.5 1.5 0 0 1 1.5-1.5Z"/>`),
		g.Group(children),
	)
}