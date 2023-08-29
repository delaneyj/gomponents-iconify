package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandRedhat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 10.5l1.436-4c.318-.876.728-1.302 1.359-1.302c.219 0 1.054.365 1.88.583c.825.219.733-.329.908-.487c.176-.158.355-.294.61-.294c.242 0 .553.048 1.692.448a20.42 20.42 0 0 1 2.204.922c1.175.582 1.426.913 1.595 1.507L18.5 12.5c2.086.898 3.5 2.357 3.5 3.682C22 17.867 20.8 20 16.043 20C9.837 20 2 15.958 2 12.68c0-1.044 1.333-1.77 4-2.18z"/><path d="M6 10.5c0 .969 4.39 3.5 9.5 3.5c1.314 0 3 .063 3-1.5"/></g>`),
		g.Group(children),
	)
}