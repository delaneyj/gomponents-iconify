package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 15a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke="currentColor" stroke-width="1.5" d="M2 15c0-1.886 0-2.828.586-3.414C3.172 11 4.114 11 6 11h12c1.886 0 2.828 0 3.414.586C22 12.172 22 13.114 22 15c0 1.886 0 2.828-.586 3.414C20.828 19 19.886 19 18 19H6c-1.886 0-2.828 0-3.414-.586C2 17.828 2 16.886 2 15Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 11L3 4m14 7l4-7m-7 11h4m-.833-9.603A5.502 5.502 0 0 0 7 5.397"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M14.965 6.658a3.001 3.001 0 0 0-5.76 0"/><path fill="currentColor" d="M13.084 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}