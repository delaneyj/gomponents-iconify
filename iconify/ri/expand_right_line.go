package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandRightLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.172 11l-4.657-4.657l1.414-1.414L21 12l-7.07 7.071l-1.415-1.414L17.172 13H8v-2h9.172ZM4 19V5h2v14H4Z"/>`),
		g.Group(children),
	)
}