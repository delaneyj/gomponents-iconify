package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 10.274a2.726 2.726 0 0 0-2.078-2.648A3.72 3.72 0 0 0 10.205 4a3.712 3.712 0 0 0-2.92 1.418a2.09 2.09 0 0 0-3.719 1.573a3.028 3.028 0 0 0-3.567 2.98A3.028 3.028 0 0 0 3.026 13H13.28a2.725 2.725 0 0 0 2.719-2.726z"/>`),
		g.Group(children),
	)
}