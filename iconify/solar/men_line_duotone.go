package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="10" cy="14" r="8" stroke="currentColor" stroke-width="1.5"/><path fill="currentColor" d="M22 2h.75a.75.75 0 0 0-.75-.75V2Zm-.75 5a.75.75 0 0 0 1.5 0h-1.5ZM17 1.25a.75.75 0 0 0 0 1.5v-1.5Zm-.97 7.78l6.5-6.5l-1.06-1.06l-6.5 6.5l1.06 1.06ZM21.25 2v5h1.5V2h-1.5ZM17 2.75h5v-1.5h-5v1.5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}