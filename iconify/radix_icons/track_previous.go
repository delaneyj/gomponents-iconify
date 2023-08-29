package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrackPrevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.95 2.75a.55.55 0 1 1 1.1 0v4.533a.5.5 0 0 1 .217-.225l9-4.75A.5.5 0 0 1 13 2.75v9.5a.5.5 0 0 1-.733.442l-9-4.75a.5.5 0 0 1-.217-.225v4.533a.55.55 0 0 1-1.1 0v-9.5ZM4.57 7.5L12 11.42V3.58L4.571 7.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}