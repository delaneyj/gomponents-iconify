package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaWideAngle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 64Q138 64 58 78q-15 57-15 114t15 114q80 14 155.5 14T369 306q15-57 15-114T369 78q-80-14-155.5-14zm-.5-43q83 0 170 16l20 3l5 19q19 67 19 133t-19 133l-5 19l-20 3q-87 16-170 16T44 347l-20-3l-5-19Q0 258 0 192T19 59l5-19l20-3q87-16 169-16z"/>`),
		g.Group(children),
	)
}