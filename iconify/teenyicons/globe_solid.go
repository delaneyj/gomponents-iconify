package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 5.5A6.486 6.486 0 0 0 11.8.625l-.662.75A5.5 5.5 0 1 1 3.834 9.6l-.667.745A6.476 6.476 0 0 0 7 11.98V14H4v1h7v-1H8v-2.019A6.5 6.5 0 0 0 14 5.5Z"/><path fill="currentColor" d="M7.5 2a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Z"/>`),
		g.Group(children),
	)
}