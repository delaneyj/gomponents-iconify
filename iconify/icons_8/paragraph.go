package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paragraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 5c-3.302 0-6 2.698-6 6s2.698 6 6 6h4v10h2V7h2v20h2V7h2V5H12zm0 2h4v8h-4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4z"/>`),
		g.Group(children),
	)
}