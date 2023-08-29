package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 0h149v64h-85v235h-1q-4 36-31 60.5T96 384q-40 0-68-28T0 288t28-68t68-28q15 0 32 6V0z"/>`),
		g.Group(children),
	)
}