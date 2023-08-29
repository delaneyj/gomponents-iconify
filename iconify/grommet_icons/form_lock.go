package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M9 11V8a3 3 0 0 1 6 0v3m-3 2v3m5 2v-7H7v7h10Z"/>`),
		g.Group(children),
	)
}