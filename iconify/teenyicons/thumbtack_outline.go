package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbtackOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 14.5L5 10M.5 5.5l9 9m-1-14l6 6m-13 0l8-5m-1 12l5-8"/>`),
		g.Group(children),
	)
}