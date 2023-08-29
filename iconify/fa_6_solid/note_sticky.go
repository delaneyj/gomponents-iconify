package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteSticky(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h224V368c0-26.5 21.5-48 48-48h112V96c0-35.3-28.7-64-64-64H64zm384 320H336c-8.8 0-16 7.2-16 16v112l32-32l64-64l32-32z"/>`),
		g.Group(children),
	)
}