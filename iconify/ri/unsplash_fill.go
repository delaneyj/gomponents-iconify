package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsplashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.501 11v5h7v-5h5.5v10h-18V11h5.5Zm7-8v5h-7V3h7Z"/>`),
		g.Group(children),
	)
}