package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Winbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.896 36.026L4.5 11.974h7.303l8.185 24.083Zm8.092.031l4.137-12.228l-3.857-11.855l-4.105 12.197m4.105-12.197l7.557-.031l8.248 24.083h-7.744l-4.204-12.197m11.948 12.197L43.5 11.943h-7.24l-4.106 12.01"/>`),
		g.Group(children),
	)
}