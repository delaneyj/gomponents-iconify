package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DimButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#eda454" d="M29 51.1h6V58h-6zM29 6h6v6.9h-6zm22.1 23H58v6h-6.9zM6 29h6.9v6H6zm5.459 19.235l4.879-4.879L20.58 47.6l-4.879 4.879zM43.42 16.401l4.878-4.879l4.243 4.243l-4.88 4.879zm-.066 31.259l4.242-4.243l4.88 4.879l-4.243 4.243zM11.522 15.702l4.243-4.243l4.879 4.88l-4.243 4.242zM32 17c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15m0 25c-5.5 0-10-4.5-10-10s4.5-10 10-10s10 4.5 10 10s-4.5 10-10 10"/>`),
		g.Group(children),
	)
}