package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.28 2.28L19 11.595V3h-2v12h12v-2h-8.594l9.313-9.28l-1.44-1.44zM3 17v2h8.594L2.28 28.28l1.44 1.44L13 20.405V29h2V17H3z"/>`),
		g.Group(children),
	)
}