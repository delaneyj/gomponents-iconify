package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.706 27.585a5.261 5.261 0 0 1-3.723-8.983l1.415 1.414a3.264 3.264 0 1 0 4.616 4.616l6.03-6.03a3.264 3.264 0 0 0-4.616-4.616l-1.414-1.414a5.264 5.264 0 0 1 7.444 7.444l-6.03 6.03a5.246 5.246 0 0 1-3.722 1.539Z"/><path fill="currentColor" d="M10.264 29.997a5.262 5.262 0 0 1-3.722-8.983l6.03-6.03a5.264 5.264 0 1 1 7.444 7.443l-1.414-1.414a3.264 3.264 0 1 0-4.616-4.615l-6.03 6.03a3.264 3.264 0 0 0 4.616 4.616l1.414 1.414a5.245 5.245 0 0 1-3.722 1.54zM2 10h8v2H2zm0-4h12v2H2zm0-4h12v2H2z"/>`),
		g.Group(children),
	)
}