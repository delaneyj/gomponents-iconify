package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardSearchBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M14 4c3.771 0 5.657 0 6.828 1.172C22 6.343 22 8.229 22 12v1M10 4C6.229 4 4.343 4 3.172 5.172C2 6.343 2 8.229 2 12c0 3.771 0 5.657 1.172 6.828C4.343 20 6.229 20 10 20h3m-3-4H6"/><circle cx="18" cy="17" r="3"/><path stroke-linecap="round" d="m20.5 19.5l1 1M2 10h5m15 0H11"/></g>`),
		g.Group(children),
	)
}