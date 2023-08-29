package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 3.881V0l6 6l-6 6V8.034C2.02 7.87 2.319 12.781 4.096 16C-.29 11.259.641 3.663 9 3.881z"/>`),
		g.Group(children),
	)
}