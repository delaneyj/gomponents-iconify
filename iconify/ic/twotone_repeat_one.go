package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneRepeatOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15V9h-1l-2 1v1h1.5v4zm6-2h-2v4H7v-3l-4 4l4 4v-3h12zM17 2v3H5v6h2V7h10v3l4-4z"/>`),
		g.Group(children),
	)
}