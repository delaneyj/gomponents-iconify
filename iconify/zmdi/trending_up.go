package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 64h128v128l-49-49l-135 134l-85-85L30 320L0 290l158-158l85 85l105-104z"/>`),
		g.Group(children),
	)
}