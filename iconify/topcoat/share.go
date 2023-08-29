package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M30.5 18.5v6l10-9.929L30.5 4.5v5c-15.3.1-15 15-15 15s5.45-7.49 15-6zm-8-13h-19c-2.46 0-3 .7-3 3v24c0 2.49.6 3 3 3h29c2.41 0 3-.451 3-3v-8l-5 4.289V30.5h-25v-20h12l5-5z"/>`),
		g.Group(children),
	)
}