package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 90.83V88a78 78 0 0 0-156 0v2.83A22 22 0 0 0 56 134h4.52l62.27 109a6 6 0 0 0 10.42 0l62.27-109H200a22 22 0 0 0 6-43.17Zm-78 137.08L74.34 134h23.89L140 207Zm8-93.91l22.85 40l-11.94 20.91L112.05 134Zm29.76 27.91L149.77 134h31.89ZM200 122H56a10 10 0 0 1 0-20a6 6 0 0 0 6-6v-8a66 66 0 0 1 132 0v8a6 6 0 0 0 6 6a10 10 0 0 1 0 20Z"/>`),
		g.Group(children),
	)
}