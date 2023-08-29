package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpressionlessFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><path fill="#664e27" d="M40 48H24c-1.5 0-1.5-4 0-4h16c1.5 0 1.5 4 0 4M27.1 32h-16c-1.5 0-1.5-4 0-4h16c1.5 0 1.5 4 0 4m25.8 0h-16c-1.5 0-1.5-4 0-4h16c1.5 0 1.5 4 0 4"/>`),
		g.Group(children),
	)
}