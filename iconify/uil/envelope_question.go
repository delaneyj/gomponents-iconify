package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeQuestion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.77 5.37A1 1 0 0 0 18.13 5a1 1 0 0 1 .87-.5a1 1 0 0 1 0 2a1 1 0 0 0 0 2A3 3 0 1 0 16.4 4a1 1 0 0 0 .37 1.37ZM21 13.5a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.91l5.88 5.89a3 3 0 0 0 4.24 0l1.64-1.64a1 1 0 1 0-1.42-1.42l-1.64 1.64a1 1 0 0 1-1.4 0L5.41 7.5H13a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1Zm-2.71-3.71a1 1 0 0 0 0 1.42l.15.12a.76.76 0 0 0 .18.09a.64.64 0 0 0 .18.06h.2a1 1 0 0 0 .71-1.71a1 1 0 0 0-1.42.02Z"/>`),
		g.Group(children),
	)
}