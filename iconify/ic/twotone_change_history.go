package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneChangeHistory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.77L5.61 18h12.78z" opacity=".3"/><path fill="currentColor" d="M12 4L2 20h20L12 4zm0 3.77L18.39 18H5.61L12 7.77z"/>`),
		g.Group(children),
	)
}