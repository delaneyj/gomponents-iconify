package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaffledTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 6.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0z"/><path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zM4 6.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 4 6.5zm6 5.5H6v-1h4v1zm.5-4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 10.5 8z"/>`),
		g.Group(children),
	)
}