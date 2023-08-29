package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paragraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2v11a1 1 0 0 1-2 0V2H8v11a1 1 0 0 1-2 0V8H4a4 4 0 1 1 0-8h9a1 1 0 0 1 0 2h-1ZM6 6V2H4a2 2 0 1 0 0 4h2Z"/>`),
		g.Group(children),
	)
}