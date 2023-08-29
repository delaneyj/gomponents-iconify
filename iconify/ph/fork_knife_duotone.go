package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkKnifeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 40v128h-56s0-104 56-128Z" opacity=".2"/><path d="M72 88V40a8 8 0 0 1 16 0v48a8 8 0 0 1-16 0Zm144-48v184a8 8 0 0 1-16 0v-48h-48a8 8 0 0 1-8-8a268.75 268.75 0 0 1 7.22-56.88c9.78-40.49 28.32-67.63 53.63-78.47A8 8 0 0 1 216 40Zm-16 13.9c-32.17 24.57-38.47 84.42-39.7 106.1H200Zm-80.11-15.21a8 8 0 1 0-15.78 2.63L112 88.63a32 32 0 0 1-64 0l7.88-47.31a8 8 0 1 0-15.78-2.63l-8 48A8.17 8.17 0 0 0 32 88a48.07 48.07 0 0 0 40 47.32V224a8 8 0 0 0 16 0v-88.68A48.07 48.07 0 0 0 128 88a8.17 8.17 0 0 0-.11-1.31Z"/></g>`),
		g.Group(children),
	)
}