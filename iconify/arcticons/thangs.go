package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thangs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="9.919" cy="23.964" r="2.77"/><path d="M9.656 17.844h14.159A6.142 6.142 0 0 1 29.97 24v6.156H9.656C6.246 30.156 3.5 27.41 3.5 24s2.745-6.156 6.156-6.156Z"/><circle cx="23.677" cy="23.964" r="2.77"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="38.247" cy="24" r="2.638"/><path d="M44.5 24c0 3.4-2.756 6.156-6.156 6.156h-6.156V24A6.156 6.156 0 0 1 44.5 24Z"/></g>`),
		g.Group(children),
	)
}