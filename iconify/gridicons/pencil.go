package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 6l5 5l-9.507 9.507a1.766 1.766 0 0 1-.012-2.485l-.003-.003a1.763 1.763 0 0 1-2.521-2.467l-.008-.008a1.765 1.765 0 0 1-2.455-.036L13 6zm7.586-.414l-2.172-2.172a2 2 0 0 0-2.828 0L14 5l5 5l1.586-1.586a2 2 0 0 0 0-2.828zM3 18v3h3a3 3 0 0 0-3-3z"/>`),
		g.Group(children),
	)
}