package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncreaseTheScale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M24 9a5 5 0 0 0-10 0v10a5 5 0 0 0 10 0V9Zm18 0a5 5 0 0 0-10 0v10a5 5 0 0 0 10 0V9Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m36 32l6 6l-6 6"/><path stroke-linecap="round" d="M6 24h1"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 38H19"/></g>`),
		g.Group(children),
	)
}