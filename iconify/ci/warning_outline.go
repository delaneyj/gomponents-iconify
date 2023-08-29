package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.266 20.998H2.733a1 1 0 0 1-.866-1.5l9.266-16a1 1 0 0 1 1.73 0l9.267 16a1 1 0 0 1-.865 1.5ZM12 5.998l-7.531 13h15.064L12 5.998Zm.995 9.001h-2V9.998h2v5.001Z"/><path fill="currentColor" d="M11 16h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}