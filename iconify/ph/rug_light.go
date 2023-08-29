package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RugLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 18a6 6 0 0 0-6 6v18h-36V24a6 6 0 0 0-12 0v18h-36V24a6 6 0 0 0-12 0v18H62V24a6 6 0 0 0-12 0v208a6 6 0 0 0 12 0v-18h36v18a6 6 0 0 0 12 0v-18h36v18a6 6 0 0 0 12 0v-18h36v18a6 6 0 0 0 12 0V24a6 6 0 0 0-6-6ZM62 54h132v148H62Zm66 120a6 6 0 0 0 5.14-2.91l24-40a6 6 0 0 0 0-6.18l-24-40a6 6 0 0 0-10.28 0l-24 40a6 6 0 0 0 0 6.18l24 40A6 6 0 0 0 128 174Zm0-74.34L145 128l-17 28.34L111 128Z"/>`),
		g.Group(children),
	)
}