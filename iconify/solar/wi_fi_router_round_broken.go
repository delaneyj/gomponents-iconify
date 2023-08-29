package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterRoundBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M17 10H7a5 5 0 0 0 0 10h10a5 5 0 0 0 4-8"/><path fill="currentColor" d="M18.33 22.335a.75.75 0 1 0 1.34-.67l-1.34.67Zm-1-2l1 2l1.34-.67l-1-2l-1.34.67Zm-11.66 2a.75.75 0 1 1-1.34-.67l1.34.67Zm1-2l-1 2l-1.34-.67l1-2l1.34.67Z"/><path stroke="currentColor" stroke-width="1.5" d="M8.5 15a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 15h6.5m3.083-9.603a5.502 5.502 0 0 0-10.167 0"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M19.38 6.658a3.001 3.001 0 0 0-5.76 0"/><path fill="currentColor" d="M17.5 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}