package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmbedVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 4H3c-1.1 0-2 .9-2 2v8c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm-7 6.5L8 9.1V11H3V6h5v1.8l2-1.3v4zm4 0L12.5 12l1.5 1.5V15l-3-3l3-3v1.5zm1 4.5v-1.5l1.5-1.5l-1.5-1.5V9l3 3l-3 3z"/>`),
		g.Group(children),
	)
}