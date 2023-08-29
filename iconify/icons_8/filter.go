package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 4v2.344l.22.28l7.78 9.72V28l1.594-1.188l4-3L19 23.5v-7.156l7.78-9.72l.22-.28V4H5zm2.28 2h17.44l-7.19 9h-3.06L7.28 6zM15 17h2v5.5L15 24v-7z"/>`),
		g.Group(children),
	)
}