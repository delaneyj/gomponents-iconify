package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBoth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path d="M0 8l6-5v3h8V3l6 5l-6 5v-3H6v3L0 8z" fill="currentColor"/>`),
		g.Group(children),
	)
}