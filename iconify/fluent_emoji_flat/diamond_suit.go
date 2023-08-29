package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#F8312F" d="M14.378 2.772L4.523 14.555a2.27 2.27 0 0 0 0 2.898l9.855 11.782a2.102 2.102 0 0 0 3.244 0l9.855-11.782a2.27 2.27 0 0 0 0-2.898L17.622 2.772a2.09 2.09 0 0 0-3.244 0Z"/>`),
		g.Group(children),
	)
}