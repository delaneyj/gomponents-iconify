package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrackNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.05 2.75a.55.55 0 1 0-1.1 0v4.533a.5.5 0 0 0-.217-.225l-9-4.75A.5.5 0 0 0 2 2.75v9.5a.5.5 0 0 0 .733.442l9-4.75a.5.5 0 0 0 .217-.225v4.533a.55.55 0 0 0 1.1 0v-9.5ZM3 11.42V3.58l7.429 3.92L3 11.42Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}