package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M21.583 7.397a5.502 5.502 0 0 0-10.167 0"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M19.38 8.658a3.001 3.001 0 0 0-5.76 0"/><path fill="currentColor" d="M7 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M13 20H6c-1.886 0-2.828 0-3.414-.586C2 18.828 2 17.886 2 16c0-1.886 0-2.828.586-3.414C3.172 12 4.114 12 6 12h12c1.886 0 2.828 0 3.414.586C22 13.172 22 14.114 22 16c0 1.886 0 2.828-.586 3.414C20.828 20 19.886 20 18 20h-1m-.5-8V9"/></g>`),
		g.Group(children),
	)
}