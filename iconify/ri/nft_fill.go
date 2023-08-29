package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm3-11l9.5 5.5v11L12 23l-9.5-5.5v-11L12 1ZM4.5 7.653v8.694l2.372 1.373l8.073-5.92l4.555 2.734v-6.88L12 3.31L4.5 7.653Z"/>`),
		g.Group(children),
	)
}