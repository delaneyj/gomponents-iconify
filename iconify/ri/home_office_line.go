package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeOfficeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.673 1.612L20.8 9h-2.973L12 3.703L6 9.158V19h5v2H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0ZM14 11h9v7h-9v-7Zm2 2v3h5v-3h-5Zm8 8H13v-2h11v2Z"/>`),
		g.Group(children),
	)
}