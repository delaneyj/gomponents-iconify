package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 467h480v32H16zm0-320h480V19H16Zm32-96h416v64H48Zm208.077 160.777L96 379.079V419h320v-37.86ZM132.709 387l123.214-128.776L377.522 387Z"/>`),
		g.Group(children),
	)
}