package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnakeZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="m35.786 39.083l2.828-2.828a6 6 0 0 0 0-8.486v0a6 6 0 0 0-8.485 0l-2.462 2.462m-10.266-6.705l7.071-7.07a6 6 0 0 0 0-8.486v0a6 6 0 0 0-8.485 0l-7.071 7.071"/><path d="m18.283 22.645l-8.66 8.66a6 6 0 0 0 0 8.486v0a6 6 0 0 0 8.485 0l9.9-9.9"/><path stroke-linejoin="round" d="M15 9H7"/></g>`),
		g.Group(children),
	)
}