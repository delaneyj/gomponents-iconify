package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xvg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.61 10.335H8L15.951 27L24 10.335h-1.59l-6.36 13.33l-6.438-13.33zM16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM9.61 10.335h12.798L24 7H8l1.61 3.335z"/>`),
		g.Group(children),
	)
}