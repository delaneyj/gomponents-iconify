package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColorAccentSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 11.167c0-.368.246-.667.55-.667h9.9c.304 0 .55.299.55.667v2.666c0 .368-.246.667-.55.667h-9.9c-.304 0-.55-.299-.55-.667v-2.666Z"/>`),
		g.Group(children),
	)
}