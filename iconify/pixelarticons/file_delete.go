package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22h10V8h-2V6h-2v2h-2V6h2V4h-2V2H3v12h2V4h8v6h6v10h-8v2zm-4-2H5v2H3v-2h2v-2H3v-2h2v2h2v-2h2v2H7v2zm0 0h2v2H7v-2z"/>`),
		g.Group(children),
	)
}