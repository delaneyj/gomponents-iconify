package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxExtraLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 28H6a2.002 2.002 0 0 1-2-2V9h2v17h20V9h2v17a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M19 21V9h-2v14h7v-2h-5zM16 9h-2l-2 6l-2-6H8l2.752 7L8 23h2l2-6l2 6h2l-2.755-7L16 9zM4 4h24v2H4z"/>`),
		g.Group(children),
	)
}