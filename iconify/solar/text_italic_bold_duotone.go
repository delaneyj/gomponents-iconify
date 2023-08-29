package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 1h12a1 1 0 1 1 0 2H9a1 1 0 0 1 0-2Zm-.744 20H3a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2H8.256Z" clip-rule="evenodd"/><path d="m13.656 3l-5.4 18h2.088l5.4-18h-2.088Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}