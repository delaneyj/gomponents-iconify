package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 5V2.5h-1.25V5h-2.5v1.25h2.5v2.5h1.25v-2.5H20V5zM6.25 5v2.5h3.536A3.757 3.757 0 0 1 6.25 10A3.755 3.755 0 0 1 2.5 6.25A3.755 3.755 0 0 1 6.25 2.5c.896 0 1.759.321 2.429.905L10.32 1.52A6.194 6.194 0 0 0 6.25 0A6.257 6.257 0 0 0 0 6.25a6.257 6.257 0 0 0 6.25 6.25a6.257 6.257 0 0 0 6.25-6.25V5H6.25z"/>`),
		g.Group(children),
	)
}