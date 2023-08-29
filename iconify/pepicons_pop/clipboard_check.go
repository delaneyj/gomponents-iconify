package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.675 2.5a1 1 0 0 1 1-1h6.643a1 1 0 0 1 1 1v3.875a1 1 0 0 1-1 1H6.675a1 1 0 0 1-1-1V2.5Zm2 1v1.875h4.643V3.5H7.675Z" clip-rule="evenodd"/><path d="M5 5v11h10V5h-1V3h2a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h2v2H5Z"/><path fill-rule="evenodd" d="M12.567 8.677a1 1 0 0 1 .256 1.39l-1.767 2.565a1.5 1.5 0 0 1-2.154.334L7.387 11.79a1 1 0 0 1 1.226-1.58l1.097.851l1.466-2.128a1 1 0 0 1 1.391-.256Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}