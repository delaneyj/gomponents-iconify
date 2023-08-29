package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndpointRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 24H26"/><circle cx="22" cy="24" r="3"/><path d="M42 40H22c-8.837 0-16-7.163-16-16S13.163 8 22 8h20"/></g>`),
		g.Group(children),
	)
}