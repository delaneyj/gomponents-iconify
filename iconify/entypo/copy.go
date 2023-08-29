package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 0H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h5v2h2v-2H8.001v-2H10v-2H8v2H4V2h6v4h2V1a1 1 0 0 0-1-1zM8 7v1h2V6H9a1 1 0 0 0-1 1zm4 13h2v-2h-2v2zm0-12h2V6h-2v2zM8 19a1 1 0 0 0 1 1h1v-2H8v1zm9-13h-1v2h2V7a1 1 0 0 0-1-1zm-1 14h1a1 1 0 0 0 1-1v-1h-2v2zm0-8h2v-2h-2v2zm0 4h2v-2h-2v2z"/>`),
		g.Group(children),
	)
}