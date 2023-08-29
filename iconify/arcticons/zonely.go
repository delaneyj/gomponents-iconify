package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zonely(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM24 42.5v-37"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.659 5.5v1"/><path stroke-dasharray="0 0 2.059 2.059" d="M11.659 8.559v31.912"/><path d="M11.659 41.5v1"/></g>`),
		g.Group(children),
	)
}