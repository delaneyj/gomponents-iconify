package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SirenRoundedLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 22v-6a8 8 0 1 0-16 0v6" opacity=".5"/><path stroke-linecap="round" d="M14.29 11.5a4.014 4.014 0 0 1 2.21 2.21M2 22h20M12 2v3m9 1l-1.5 1.5M3 6l1.5 1.5"/><path d="M13.5 17.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path stroke-linecap="round" d="M12 19v3" opacity=".5"/></g>`),
		g.Group(children),
	)
}