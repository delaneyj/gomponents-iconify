package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tsv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28 9l-2 13l-2-13h-2l2.516 14h2.968L30 9h-2zM18 23h-6v-2h6v-4h-4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h6v2h-6v4h4a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM2 11h3v12h2V11h3V9H2v2z"/>`),
		g.Group(children),
	)
}