package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.745 5.728A6.5 6.5 0 0 1 18.5 7.5V9h2v13h-17V9h13V7.5a4.5 4.5 0 0 0-8.83-1.228l-.273.962l-1.924-.544l.272-.962ZM5.5 11v9h13v-9h-13ZM9 14.5h6v2H9v-2Z"/>`),
		g.Group(children),
	)
}