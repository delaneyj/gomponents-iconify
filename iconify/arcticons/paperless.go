package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.461 40.263Q4.626 36.386 7.143 22.47c3.343 6.384 14.023 10.155 7.318 17.793zm0 0a6.388 6.388 0 0 1 1.625 2.586"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.813 43.5a14.419 14.419 0 0 1 4.064-6.1Q44.339 31.487 40.96 4.5C35.9 15.464 10.594 17.989 19.877 37.4m-9.017-7.86a8.54 8.54 0 0 1 4.17 7.745"/>`),
		g.Group(children),
	)
}