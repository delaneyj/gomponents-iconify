package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.3 19.5v-15h2.4v15H17L12.5 24L8 19.5zM12.5 0L17 4.5H8z"/>`),
		g.Group(children),
	)
}