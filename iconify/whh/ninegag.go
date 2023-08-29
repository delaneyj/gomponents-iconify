package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ninegag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 1024L0 768l192-96l256 160l256-160V448L448 576L0 352v-96L448 0l448 256v512zm192-736l-192-96l-192 96l192 96z"/>`),
		g.Group(children),
	)
}