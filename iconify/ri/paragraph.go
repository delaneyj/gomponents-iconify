package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paragraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6v15h-2v-5a6 6 0 0 1 0-12h10v2h-3v15h-2V6h-3Zm-2 0a4 4 0 1 0 0 8V6Z"/>`),
		g.Group(children),
	)
}