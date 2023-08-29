package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 7v2h24V7zm0 8v2h24v-2zm0 8v2h24v-2z"/>`),
		g.Group(children),
	)
}