package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="#757f3f"><circle cx="64" cy="17.12" r="15.31"/><circle cx="64" cy="110.88" r="15.31"/><path d="M8.57 51.69c-4.36 0-7.92 3.56-7.92 7.92v8.78c0 4.36 3.56 7.92 7.92 7.92h110.86c4.36 0 7.92-3.56 7.92-7.92v-8.78c0-4.36-3.56-7.92-7.92-7.92H8.57z"/></g>`),
		g.Group(children),
	)
}