package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RandOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 14V8.5m0 0v-7H8a3.5 3.5 0 1 1 0 7H3.5Zm0 0h5a3 3 0 0 1 3 3V14"/>`),
		g.Group(children),
	)
}