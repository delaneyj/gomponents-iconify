package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashEco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 6l1.721 10.329A2 2 0 0 0 6.694 18H12m8.162-6.972L21 6"/><path d="M3.486 8.965c.168.02.34.033.514.035c.79.009 1.539-.178 2-.5c.461-.32 1.21-.507 2-.5c.79-.007 1.539.18 2 .5c.461.322 1.21.509 2 .5c.79.009 1.539-.178 2-.5c.461-.32 1.21-.507 2-.5c.79-.007 1.539.18 2 .5c.461.322 1.21.509 2 .5c.17 0 .339-.014.503-.034M16 22s0-2 3-4"/><path d="M19 21a3 3 0 0 1 0-6h3v3a3 3 0 0 1-3 3z"/></g>`),
		g.Group(children),
	)
}