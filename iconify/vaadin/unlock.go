package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 8V4.9C8 2.7 6.2 1 4.1 1h-.3C1.6 1 0 2.7 0 4.9V7h2V4.9C2 3.8 2.7 3 3.8 3h.3c1 0 1.9.8 1.9 1.9V8H5l.1 5S5 16 10 16s5-3 5-3V8H8zm3 6h-1v-1.8c-.6 0-1-.6-1-1.1c0-.6.4-1.1 1-1.1s1 .4 1 .9V14z"/>`),
		g.Group(children),
	)
}