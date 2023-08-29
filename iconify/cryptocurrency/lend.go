package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm1.147-11.554l-5.474-5.7L10 16.45l3.92 3.995l-3.777 3.849L11.697 26zm3.18-3.191L22 15.549l-3.92-3.995l3.777-3.849L20.303 6l-5.474 5.554zm-7.96-3.167l5.498 5.7l1.673-1.705l-5.498-5.603z"/>`),
		g.Group(children),
	)
}