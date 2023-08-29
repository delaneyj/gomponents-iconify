package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 3.578a.75.75 0 0 1 0-1.5h3.536a.75.75 0 0 1 0 1.5H10Zm-3.47.452a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M12 5.06a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17Zm4.99 3.711a.5.5 0 0 0-.7-.701l-3.175 2.468l-2.075 1.483a1.434 1.434 0 1 0 2 2l1.482-2.076l2.469-3.174Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}