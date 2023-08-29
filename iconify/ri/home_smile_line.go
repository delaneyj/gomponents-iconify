package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSmileLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19h12V9.158l-6-5.455l-6 5.455V19Zm13 2H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0L23 11h-3v9a1 1 0 0 1-1 1ZM7.5 13h2a2.5 2.5 0 0 0 5 0h2a4.5 4.5 0 1 1-9 0Z"/>`),
		g.Group(children),
	)
}