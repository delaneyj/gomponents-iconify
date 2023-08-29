package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UngroupObjects(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 14v-4h-4v2h-6V6h2V2h-4v2H6V2H2v4h2v12H2v4h4v-2h6v6h-2v4h4v-2h12v2h4v-4h-2V14ZM6 6h12v12H6Zm20 20H14v-6h4v2h4v-4h-2v-4h6Z"/>`),
		g.Group(children),
	)
}