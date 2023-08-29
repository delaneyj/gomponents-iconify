package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBar0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBar1" fill="currentColor"><path id="feBar2" d="M3 16h18v2H3v-2Zm0-5h18v2H3v-2Zm0-5h18v2H3V6Z"/></g></g>`),
		g.Group(children),
	)
}