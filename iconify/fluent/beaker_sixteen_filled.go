package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3v3.689a2.5 2.5 0 0 1-.34 1.26L5.045 9h5.908l-.614-1.051a2.5 2.5 0 0 1-.34-1.26V3h.5a.5.5 0 0 0 0-1h-5a.5.5 0 0 0 0 1H6Zm5.537 7H4.463l-1.018 1.744A1.5 1.5 0 0 0 4.741 14h6.518a1.5 1.5 0 0 0 1.296-2.256L11.537 10Z"/>`),
		g.Group(children),
	)
}