package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nmr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" d="M32.485 15.515A11.962 11.962 0 0 0 24 12a11.962 11.962 0 0 0-8.485 3.515"/><path stroke-linejoin="round" d="M34 24H14v8h20v-8Z"/><path stroke-linecap="round" d="M17.045 32L17 42.715M31.046 32L31 42.715"/></g>`),
		g.Group(children),
	)
}