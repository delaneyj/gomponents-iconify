package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 7.5a6.5 6.5 0 0 1 13 0V9h2v2.5h-2V11h-13v9h6v2h-8V9h2V7.5Zm2 1.5h9V7.5a4.5 4.5 0 1 0-9 0V9Zm11 6a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM13 18.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm-4-4h2.5v2H9v-2Zm10.5 1.752v1.834l1.414 1.414l-1.414 1.414l-2-2v-2.662h2Z"/>`),
		g.Group(children),
	)
}