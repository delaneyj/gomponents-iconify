package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Farm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 20H2V8l6-3l6 3v12h-3m-6 0v-7h6v7m-6 0h6m7-6v6m-4-3h8m-8-3h8"/>`),
		g.Group(children),
	)
}