package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsShrinkV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 1v2H7V1h10Zm-.757 7.448l-1.414 1.414L13 8.033v7.934l1.829-1.829l1.414 1.414L12 19.795l-4.243-4.243l1.415-1.414L11 15.966V8.034L9.172 9.862L7.757 8.448L12 4.205l4.243 4.243ZM17 23v-2H7v2h10Z"/>`),
		g.Group(children),
	)
}