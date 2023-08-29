package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ptv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 34.073V18.754h5.3c2.997 0 5.426 2.304 5.426 5.145s-2.429 5.144-5.426 5.144H4.5m12.396-10.56h10.746m-5.373 15.866V18.483m5.364-4.833L43.5 19.023l-15.867 5.373"/>`),
		g.Group(children),
	)
}