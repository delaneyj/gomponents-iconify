package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoardSplitFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.25 6A5.25 5.25 0 0 0 6 11.25V23h21V6H11.25ZM27 25.5H6v11.25C6 39.65 8.35 42 11.25 42H27V25.5ZM36.75 42H29.5V31.5H42v5.25c0 2.9-2.35 5.25-5.25 5.25ZM42 19.5V29H29.5v-9.5H42Zm0-2.5H29.5V6h7.25C39.65 6 42 8.35 42 11.25V17Z"/>`),
		g.Group(children),
	)
}