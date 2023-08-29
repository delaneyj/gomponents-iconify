package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOutDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m216 128l-40 40V88Z" opacity=".2"/><path d="M112 216a8 8 0 0 1-8 8H48a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h56a8 8 0 0 1 0 16H48v160h56a8 8 0 0 1 8 8Zm109.66-82.34l-40 40A8 8 0 0 1 168 168v-32h-64a8 8 0 0 1 0-16h64V88a8 8 0 0 1 13.66-5.66l40 40a8 8 0 0 1 0 11.32Zm-17-5.66L184 107.31v41.38Z"/></g>`),
		g.Group(children),
	)
}