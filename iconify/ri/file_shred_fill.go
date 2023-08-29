package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileShredFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12v2H2v-2h2V2.995c0-.55.445-.995.996-.995H15l5 5v5h2ZM3 16h2v6H3v-6Zm16 0h2v6h-2v-6Zm-4 0h2v6h-2v-6Zm-4 0h2v6h-2v-6Zm-4 0h2v6H7v-6Z"/>`),
		g.Group(children),
	)
}