package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordCircleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z" opacity=".5"/><path fill-rule="evenodd" d="M12.75 12c0 .644.188 1.245.51 1.75h-2.52A3.25 3.25 0 1 0 8 15.25h8A3.25 3.25 0 1 0 12.75 12Zm1.5 0a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0Zm-4.5 0a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}