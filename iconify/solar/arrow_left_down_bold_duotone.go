package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftDownBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m11.56 13.5l-1.06-1.06l-3.97-3.97A.75.75 0 0 0 5.25 9v9c0 .414.336.75.75.75h9a.75.75 0 0 0 .53-1.28l-3.97-3.97Z" clip-rule="evenodd"/><path d="M18.53 6.53a.75.75 0 0 0-1.06-1.06l-6.97 6.97l1.06 1.06l6.97-6.97Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}