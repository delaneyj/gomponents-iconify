package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shrink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.672 10.561a.7.7 0 0 1 .7.7v6a.7.7 0 1 1-1.4 0l-.001-4.354l-6.772 6.889a.7.7 0 1 1-.998-.982l6.737-6.854l-4.266.001a.7.7 0 1 1 0-1.4h6ZM19.792.201a.7.7 0 0 1 .009.99l-6.737 6.854l4.266-.001a.7.7 0 0 1 0 1.4h-6a.7.7 0 0 1-.7-.7v-6a.7.7 0 1 1 1.4 0v4.354L18.804.209a.7.7 0 0 1 .99-.008Z"/>`),
		g.Group(children),
	)
}