package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ardr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-.117-14.81L12.5 23l5.152-3.498l-1.769-2.312zM16 6L6 23h5.455l7.272-12.526L16 6zm0 9.842L21.455 23H26l-6.364-9.842L16 15.842z"/>`),
		g.Group(children),
	)
}