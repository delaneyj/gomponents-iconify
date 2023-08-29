package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m34 10l8 8m0-8l-8 8m10 12l-7 8l-4-4m-7-24H4v8h22v-8Zm0 20H4v8h22v-8Z"/>`),
		g.Group(children),
	)
}