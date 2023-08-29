package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drumstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.45 15.4c-2.13.65-4.3.32-5.7-1.1c-2.29-2.27-1.76-6.5 1.17-9.42c2.93-2.93 7.15-3.46 9.43-1.18c1.41 1.41 1.74 3.57 1.1 5.71c-1.4-.51-3.26-.02-4.64 1.36c-1.38 1.38-1.87 3.23-1.36 4.63z"/><path d="m11.25 15.6l-2.16 2.16a2.5 2.5 0 1 1-4.56 1.73a2.49 2.49 0 0 1-1.41-4.24a2.5 2.5 0 0 1 3.14-.32l2.16-2.16"/></g>`),
		g.Group(children),
	)
}