package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6 7.07A8 8 0 1 1 3.07 10"/><path fill="currentColor" d="M22 2h.75a.75.75 0 0 0-.75-.75V2Zm-.75 5a.75.75 0 0 0 1.5 0h-1.5ZM17 1.25a.75.75 0 0 0 0 1.5v-1.5Zm-.97 7.78l6.5-6.5l-1.06-1.06l-6.5 6.5l1.06 1.06ZM21.25 2v5h1.5V2h-1.5ZM17 2.75h5v-1.5h-5v1.5Z"/></g>`),
		g.Group(children),
	)
}