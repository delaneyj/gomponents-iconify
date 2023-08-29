package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 11V9h-8v14h2v-6h5v-2h-5v-4h6zm-18 0h3v10h-3v2h8v-2h-3V11h3V9h-8v2zM2 11h3v12h2V11h3V9H2v2z"/>`),
		g.Group(children),
	)
}