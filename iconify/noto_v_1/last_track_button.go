package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="#40c0e7"><path d="M121.76 32.82L72.82 61.08V32.82L18.83 64l53.99 31.18V66.92l48.94 28.26z"/><path d="M6.24 33.23h17.58v61.54H6.24z"/></g>`),
		g.Group(children),
	)
}