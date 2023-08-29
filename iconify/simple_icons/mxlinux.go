package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mxlinux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.001 13.301l3.277 3.819l-.75.9l-2.133-2.521l-1.131-1.338l.737-.86zM24 2.41v19.182c0 .655-.531 1.186-1.186 1.186H1.186A1.186 1.186 0 0 1 0 21.591V2.409c0-.655.531-1.186 1.186-1.186h21.628c.655 0 1.186.53 1.186 1.186zM21.759 19.5l-2.116-2.542l-2.115-2.541l-.586.704l-3.25-3.788l4.913-5.73l-1.175-1.008l-4.76 5.549l-4.743-5.527l-1.947 1.67l5 5.827l-.73.851L9.01 11.5l-3.384 4l-3.385 4h19.518z"/>`),
		g.Group(children),
	)
}