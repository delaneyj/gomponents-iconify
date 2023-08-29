package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandOpenSource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a9 9 0 0 1 3.618 17.243l-2.193-5.602a3 3 0 1 0-2.849 0l-2.193 5.603A9 9 0 0 1 12 3z"/>`),
		g.Group(children),
	)
}