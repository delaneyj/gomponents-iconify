package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeWithMeridians(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="28" fill="#92D3F5"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="28"/><path d="M36 8v56c-8.56 0-15.5-12.536-15.5-28S27.44 8 36 8c8.56 0 15.5 12.536 15.5 28S44.56 64 36 64m28-28H8m52-14H12m48 28H12"/></g>`),
		g.Group(children),
	)
}