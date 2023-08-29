package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Puzzle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPuzzle0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M4 24V12h9v-2a6 6 0 0 1 12 0v2h9v12h4a6 6 0 0 1 0 12h-4v8H4v-8h4a6 6 0 0 0 0-12H4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPuzzle0)"/>`),
		g.Group(children),
	)
}