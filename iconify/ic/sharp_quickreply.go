package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpQuickreply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H2v20l4-4h9v-8h7z"/><path fill="currentColor" d="M22.5 16h-2.2l1.7-4h-5v6h2v5z"/>`),
		g.Group(children),
	)
}