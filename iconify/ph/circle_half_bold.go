package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm12 24.87a83.53 83.53 0 0 1 24 7.25v151.76a83.53 83.53 0 0 1-24 7.25ZM44 128a84.12 84.12 0 0 1 72-83.13v166.26A84.12 84.12 0 0 1 44 128Zm144 58.71V69.29a83.81 83.81 0 0 1 0 117.42Z"/>`),
		g.Group(children),
	)
}