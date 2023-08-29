package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uikit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m22.2 6.7l-4.3 2.5l5.1 2.9v7.8L16 24l-7-4.1v-5.5L5 12v10.5L16 29l11-6.5v-13zm-2-1.2L16 3l-4.2 2.5L16 8z"/>`),
		g.Group(children),
	)
}