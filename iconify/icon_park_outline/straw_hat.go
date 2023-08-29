package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrawHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M34 22c0-5.523-4.477-10-10-10s-10 4.477-10 10m0 1c-5.978 1.204-10 3.456-10 6.034C4 32.881 12.954 36 24 36s20-3.119 20-6.966c0-2.578-4.022-4.83-10-6.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 22c1 .833 4 4 10 4s9-3 10-4m-15 3l2-5"/></g>`),
		g.Group(children),
	)
}