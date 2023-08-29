package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XServer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.372 4.5v31.87l6.986-6.939L26.464 43.5l6.291-4.067l-6.889-12.897h9.762L12.372 4.5z"/>`),
		g.Group(children),
	)
}