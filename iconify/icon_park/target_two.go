package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M24 41C33.39 41 41 33.39 41 24C41 14.61 33.39 7 24 7C14.61 7 7 14.61 7 24C7 33.39 14.61 41 24 41Z"/><path d="M24 32C28.42 32 32 28.42 32 24C32 19.58 28.42 16 24 16C19.58 16 16 19.58 16 24C16 28.42 19.58 32 24 32Z"/><path stroke-linecap="round" d="M4 24H44"/><path stroke-linecap="round" d="M24 4V44"/></g>`),
		g.Group(children),
	)
}