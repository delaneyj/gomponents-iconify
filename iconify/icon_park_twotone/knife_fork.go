package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTKnifeFork0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 4v40M8 5v10c0 5 6 5 6 5s6 0 6-5V5"/><path fill="#555" d="M30 12c0-8 8-8 8-8v17h-8v-9Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 21h-8v-9c0-8 8-8 8-8v17Zm0 0v23"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTKnifeFork0)"/>`),
		g.Group(children),
	)
}