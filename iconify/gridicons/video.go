package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.001 4h8v1.997h2V4A2 2 0 0 1 20 6v12a2 2 0 0 1-1.999 2v-2.003h-2V20h-8v-2.003h-2V20H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h.001v1.997h2V4zM10 15l4.5-3L10 9v6zm8.001.997v-3h-2v3h2zm0-5v-3h-2v3h2zm-10 5v-3h-2v3h2zm0-5v-3h-2v3h2z"/>`),
		g.Group(children),
	)
}