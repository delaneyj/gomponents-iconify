package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 7.5a6.5 6.5 0 0 1 13 0V9h2v13h-17V9h2V7.5Zm2 1.5h9V7.5a4.5 4.5 0 1 0-9 0V9Zm-2 2v9h13v-9h-13ZM9 14.5h6v2H9v-2Z"/>`),
		g.Group(children),
	)
}