package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.37 11.84a13.6 13.6 0 0 0-1.06-2.51A14.17 14.17 0 0 0 25.9 6.1a14 14 0 1 0 0 19.8a14.17 14.17 0 0 0 2.41-3.23a13.6 13.6 0 0 0 1.06-2.51a14 14 0 0 0 0-8.32ZM4 16A12 12 0 0 1 16 4v24A12 12 0 0 1 4 16Z"/>`),
		g.Group(children),
	)
}