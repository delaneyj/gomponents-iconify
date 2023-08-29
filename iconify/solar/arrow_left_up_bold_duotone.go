package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftUpBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.53 15.53A.75.75 0 0 1 5.25 15V6A.75.75 0 0 1 6 5.25h9a.75.75 0 0 1 .53 1.28l-9 9Z" clip-rule="evenodd"/><path d="M18.53 17.47a.75.75 0 1 1-1.06 1.06l-6.97-6.97l1.06-1.06l6.97 6.97Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}