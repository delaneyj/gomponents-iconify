package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Breather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25.192 25.192A13 13 0 1 0 6.808 6.808a13 13 0 0 0 18.384 18.384Z"/>`),
		g.Group(children),
	)
}