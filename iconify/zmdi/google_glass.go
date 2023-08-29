package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M235 195h181v21q0 89-58 151t-145 62q-88 0-150.5-62.5T0 216T62.5 65.5T213 3q89 0 148 65l-38 38q-43-50-110-50q-66 0-113 47T53 216t47 113t113 47q56 0 96.5-36t50.5-92H235v-53z"/>`),
		g.Group(children),
	)
}