package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M.5 7.5a7 7 0 1 1 14 0a7 7 0 0 1-14 0Z"/><path d="M9.5 5.5h-4v4h4v-4Z"/></g>`),
		g.Group(children),
	)
}