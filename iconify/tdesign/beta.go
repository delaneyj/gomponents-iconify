package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2.491v18.51H4v-2h14v-3.058L3.49 9.476L20 2.491Zm-2 11.263V5.509L8.51 9.524l9.49 4.23Z"/>`),
		g.Group(children),
	)
}