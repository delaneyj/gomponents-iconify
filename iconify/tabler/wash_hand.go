package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashHand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.486 8.965c.168.02.34.033.514.035c.79.009 1.539-.178 2-.5c.426-.296.777-.5 1.5-.5h1M16 8l.615.034c.552.067 1.046.23 1.385.466c.461.322 1.21.509 2 .5c.17 0 .339-.014.503-.034M14 10.5l.586.578a1.516 1.516 0 0 0 2 0c.476-.433.55-1.112.176-1.622L15 7c-.37-.506-1.331-1-2-1H9.883a1 1 0 0 0-.992.876l-.499 3.986A3.857 3.857 0 0 0 11 15a2.28 2.28 0 0 0 3-2.162V10.5z"/><path d="m3 6l1.721 10.329A2 2 0 0 0 6.694 18h10.612a2 2 0 0 0 1.973-1.671L21 6"/></g>`),
		g.Group(children),
	)
}