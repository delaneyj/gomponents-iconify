package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 6v15H8V6H2V4h14v2h-6Zm8 8v7h-2v-7h-3v-2h8v2h-3Z"/>`),
		g.Group(children),
	)
}