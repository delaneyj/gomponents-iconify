package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartAngleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 9.137c0 4.405 3.298 6.946 6.106 9.11c.292.225.579.445.856.664C10 19.729 11 20.5 12 20.5v-15C7.5.826 2 4.275 2 9.138Z" clip-rule="evenodd" opacity=".5"/><path d="m14 7.5l-2-2v15c1 0 2-.77 3.038-1.59c.277-.218.564-.438.856-.663C18.702 16.083 22 13.542 22 9.137c0-4.462-4.631-7.734-8.871-4.63l1.931 1.931A.75.75 0 0 1 14 7.5Z"/></g>`),
		g.Group(children),
	)
}