package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12.984 5.016H21v13.968h-8.016V5.016zM3 5.016h8.016v1.968H3V5.016zm0 6V9h8.016v2.016H3zm0 7.968v-1.968h8.016v1.968H3zM3 15v-2.016h8.016V15H3z" fill="currentColor"/>`),
		g.Group(children),
	)
}