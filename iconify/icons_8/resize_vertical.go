package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 2.094l-.72.687l-8 8l1.44 1.44L15 5.936v20.125l-6.28-6.28l-1.44 1.437l8 8l.72.686l.72-.687l8-8l-1.44-1.44L17 26.063V5.938l6.28 6.282l1.44-1.44l-8-8l-.72-.686z"/>`),
		g.Group(children),
	)
}