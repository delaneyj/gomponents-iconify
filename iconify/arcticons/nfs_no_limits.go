package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NfsNoLimits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.328 29.728l2.435-11.456l5.155 11.456l2.435-11.456m9.937 10.201c.508.915 1.316 1.255 2.542 1.255h1.696c1.579 0 3.13-1.28 3.466-2.858l.002-.012c.336-1.579-.672-2.858-2.25-2.858h-1.871c-1.58 0-2.59-1.281-2.253-2.861h0c.336-1.584 1.893-2.867 3.477-2.867h1.686c1.226 0 2.035.34 2.543 1.255M21.736 24h3.723m-4.941 5.728l2.436-11.456h5.728"/>`),
		g.Group(children),
	)
}