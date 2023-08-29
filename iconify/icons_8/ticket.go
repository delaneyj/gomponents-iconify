package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 7v7h1c1.19 0 2 .81 2 2s-.81 2-2 2H2v7h28v-7h-1c-1.19 0-2-.81-2-2s.81-2 2-2h1V7H2zm2 2h24v3.188c-1.715.45-3 1.955-3 3.812c0 1.857 1.285 3.362 3 3.813V23H4v-3.188c1.715-.45 3-1.955 3-3.812c0-1.857-1.285-3.362-3-3.813V9z"/>`),
		g.Group(children),
	)
}