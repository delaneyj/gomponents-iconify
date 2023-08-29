package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8v13h2V8a3.999 3.999 0 0 0-4-4H8.243l2 2H26a1.996 1.996 0 0 1 2 2zm2 20.586L3.414 2L2 3.414l1.504 1.504A3.918 3.918 0 0 0 2 8v12a4 4 0 0 0 4 4h6v-2H6a1.996 1.996 0 0 1-2-2V8a1.981 1.981 0 0 1 .92-1.667L20.585 22H17l-4 7l1.736 1l3.429-6h4.42l6 6z"/>`),
		g.Group(children),
	)
}