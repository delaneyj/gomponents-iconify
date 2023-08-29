package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="#40c0e7"><path d="m6.24 95.18l48.94-28.26v28.26L109.17 64L55.18 32.82v28.26L6.24 32.82z"/><path d="M104.18 33.23h17.58v61.54h-17.58z"/></g>`),
		g.Group(children),
	)
}