package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.657 3a1 1 0 0 1 1-1H16a3 3 0 0 1 3 3v7.343a1 1 0 0 1-2 0V5a1 1 0 0 0-1-1H8.657a1 1 0 0 1-1-1Zm11.061 14.304a.831.831 0 0 0-.023-.023L6.72 5.304a.997.997 0 0 0-.023-.023l-1.99-1.988a1 1 0 0 0-1.414 1.414L5 6.414V21a1 1 0 0 0 1.447.894L12 19.118l5.553 2.776A1 1 0 0 0 19 21v-.586l.293.293a1 1 0 0 0 1.414-1.414l-1.989-1.989Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}