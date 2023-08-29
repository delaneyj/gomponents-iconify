package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM8 9v14h1V9H8zm5 3v9h1v-9h-1zm5 2v5h1v-5h-1zm5-5v14h1V9h-1z"/>`),
		g.Group(children),
	)
}