package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paragraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 4H13a7 7 0 0 0 0 14v10h2V6h5v22h2V6h5ZM13 16a5 5 0 0 1 0-10Z"/>`),
		g.Group(children),
	)
}