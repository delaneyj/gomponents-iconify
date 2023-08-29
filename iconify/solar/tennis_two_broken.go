package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TennisTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-dasharray="1.5 2" d="M13.196 2.071s-1.098 4.099 1.402 8.43c2.5 4.33 6.598 5.428 6.598 5.428M2.803 8.071s4.099 1.099 6.599 5.43c2.5 4.33 1.402 8.428 1.402 8.428"/><path stroke-linecap="round" d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}