package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 136H48a20 20 0 0 0-20 20v36a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20v-36a20 20 0 0 0-20-20Zm-4 52H52v-28h152Zm4-144H48a20 20 0 0 0-20 20v36a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20Zm-4 52H52V68h152Z"/>`),
		g.Group(children),
	)
}