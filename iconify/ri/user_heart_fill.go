package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserHeartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.841 15.659l.176.177l.178-.177a2.25 2.25 0 1 1 3.181 3.182L18.018 22.2l-3.359-3.359a2.25 2.25 0 1 1 3.182-3.182ZM12 14v8H4a8 8 0 0 1 7.75-7.996L12 14Zm0-13c3.315 0 6 2.685 6 6s-2.685 6-6 6s-6-2.685-6-6s2.685-6 6-6Z"/>`),
		g.Group(children),
	)
}