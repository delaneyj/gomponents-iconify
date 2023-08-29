package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpWebAssetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.83 4H22v15.17l-2-2V8h-9.17l-4-4zm13.66 19.31L17.17 20H2V4.83L.69 3.51L2.1 2.1l19.8 19.8l-1.41 1.41zM15.17 18l-10-10H4v10h11.17z"/>`),
		g.Group(children),
	)
}