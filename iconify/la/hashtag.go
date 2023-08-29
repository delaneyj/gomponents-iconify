package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hashtag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 5v6H5v2h6v6H5v2h6v6h2v-6h6v6h2v-6h6v-2h-6v-6h6v-2h-6V5h-2v6h-6V5zm2 8h6v6h-6z"/>`),
		g.Group(children),
	)
}