package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WicketIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FF9925" d="M127.998 0C198.694 0 256 57.306 256 127.998c0 70.696-57.306 128.005-128.002 128.005C57.306 256.003 0 198.693 0 127.998C0 57.306 57.306 0 127.998 0"/><path fill="#FFF" d="M79.988 79.99v96.024h96.024V79.99l-32.01 64.013l-16.004-32.006l-16 32.006z"/>`),
		g.Group(children),
	)
}