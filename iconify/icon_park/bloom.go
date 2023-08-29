package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bloom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 32L18 42"/><path d="M42 32L30 42"/><path d="M24 32V44"/><path fill="#2F88FF" d="M17 11L24 4L31 11L38 10C38 10 39 14.2386 39 17C39 27 30.5 32 24 32C17.5 32 9 27 9 17C9 14.2386 10 10 10 10L17 11Z"/></g>`),
		g.Group(children),
	)
}