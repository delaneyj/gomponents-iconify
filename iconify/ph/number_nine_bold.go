package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 96a60 60 0 1 0-60 60a59.21 59.21 0 0 0 7.81-.53l-26.27 46.64a12 12 0 0 0 20.92 11.78l49.54-88A59.57 59.57 0 0 0 188 96Zm-96 0a36 36 0 1 1 36 36a36 36 0 0 1-36-36Z"/>`),
		g.Group(children),
	)
}