package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.17 13.42a5.24 5.24 0 0 1-.93-2.06L10.7 9a1 1 0 0 0 0-1.39l-3.65-4.1a1 1 0 0 0-1.4-.08L3.48 5.29a1 1 0 0 0-.29.65a15.25 15.25 0 0 0 3.54 9.92l-4.44 4.43l1.42 1.42l18-18l-1.42-1.42zm7.44.02a1 1 0 0 0-1.39.05L12.82 16a4.07 4.07 0 0 1-.51-.14l-2.66 2.61A15.46 15.46 0 0 0 17.89 21h.36a1 1 0 0 0 .65-.29l1.86-2.17a1 1 0 0 0-.09-1.39z"/>`),
		g.Group(children),
	)
}