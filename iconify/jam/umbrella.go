package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 3a8.003 8.003 0 0 0-7.748 6h15.496C16.86 5.55 13.728 3 10 3zm-1 8H0c0-5.186 3.947-9.45 9.001-9.95L9 1a1 1 0 1 1 1.999.05C16.053 1.55 20 5.813 20 11h-9v6a1 1 0 0 0 2 0a1 1 0 0 1 2 0a3 3 0 0 1-6 0v-6z"/>`),
		g.Group(children),
	)
}