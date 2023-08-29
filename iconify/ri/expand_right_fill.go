package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.586 4.586L21 12l-7.414 7.414V13H8v-2h5.586V4.586ZM4 19V5h2v14H4Z"/>`),
		g.Group(children),
	)
}