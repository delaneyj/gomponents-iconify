package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisHorizontalCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm8 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-3-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm7 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}