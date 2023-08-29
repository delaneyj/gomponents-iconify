package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6v3.586l-3.657-3.657c-1.172-1.172-2.707-1.757-4.243-1.757s-3.071.585-4.242 1.757a6 6 0 0 0 0 8.485l5.364 5.364l1.414-1.414L7.272 13a4.004 4.004 0 0 1 0-5.657A3.975 3.975 0 0 1 10.1 6.171c1.068 0 2.073.416 2.828 1.172L16.586 11H13v2h7V6h-2z"/>`),
		g.Group(children),
	)
}