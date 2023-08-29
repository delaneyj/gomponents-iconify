package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkullSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 7.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm6 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v3.382a1.5 1.5 0 0 1-.83 1.342l-1.17.585v.691A2.5 2.5 0 0 1 9.5 15h-4A2.5 2.5 0 0 1 3 12.5v-.691l-1.17-.585A1.5 1.5 0 0 1 1 9.882V6.5ZM4.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm6 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM7.146 9.146a.5.5 0 0 1 .708 0l1 1l-.708.708l-.646-.647l-.646.647l-.708-.708l1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}