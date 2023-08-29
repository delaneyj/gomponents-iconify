package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightDownBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.47 8.47a.75.75 0 0 1 1.28.53v9a.75.75 0 0 1-.75.75H9a.75.75 0 0 1-.53-1.28l9-9Z" clip-rule="evenodd"/><path d="M5.47 6.53a.75.75 0 0 1 1.06-1.06l6.97 6.97l-1.06 1.06l-6.97-6.97Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}