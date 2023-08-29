package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2v4H3V2h18M4 7h16v13H4V7m5 5h6v-1.5c0-.28-.22-.5-.5-.5h-5c-.28 0-.5.22-.5.5V12M7 24h2v-2H7v2m4 0h2v-2h-2v2m4 0h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}