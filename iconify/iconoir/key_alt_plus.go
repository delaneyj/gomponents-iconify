package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyAltPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.992 18h3m3 0h-3m0 0v-3m0 3v3m-5.58-10.657a4 4 0 1 0 5.657-5.657a4 4 0 0 0-5.657 5.657Zm0 0l-8.485 8.485l2.121 2.122M6.755 16l2.122 2.121"/>`),
		g.Group(children),
	)
}