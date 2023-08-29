package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortHorizontalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.663 3.478a.75.75 0 0 1-.016 1.06L7.851 7.25H18a.75.75 0 0 1 0 1.5H7.85l2.797 2.712a.75.75 0 1 1-1.044 1.076l-4.125-4a.75.75 0 0 1 0-1.076l4.125-4a.75.75 0 0 1 1.06.016Zm2.674 8a.75.75 0 0 1 1.06-.016l4.125 4a.75.75 0 0 1 0 1.076l-4.125 4a.75.75 0 1 1-1.044-1.076l2.796-2.712H6a.75.75 0 0 1 0-1.5h10.15l-2.797-2.712a.75.75 0 0 1-.016-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}