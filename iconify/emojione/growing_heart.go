package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrowingHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#ff5a79"><path d="M61.1 18.2c-6.4-17-27.2-9.4-29.1-.9c-2.6-9-22.9-15.7-29.1.9C-4 36.7 29.6 53.3 32 56c2.4-2.2 36-19.6 29.1-37.8M32 50.8C30 48.6 2 35.6 7.8 21c5.1-13.1 22-7.8 24.2-.7c1.6-6.6 18.9-12.6 24.2.7C62 35.3 34 49.1 32 50.8"/><path d="M49.2 24.8c-3.8-10-16.1-5.5-17.2-.5c-1.6-5.3-13.6-9.3-17.2.5c-4.1 11 15.8 20.8 17.2 22.4c1.4-1.3 21.3-11.6 17.2-22.4"/></g>`),
		g.Group(children),
	)
}