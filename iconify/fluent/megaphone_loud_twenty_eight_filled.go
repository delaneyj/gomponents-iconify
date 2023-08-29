package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneLoudTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.525 4.466a.75.75 0 1 0 1.45.388l.646-2.415a.75.75 0 1 0-1.448-.388l-.648 2.415ZM24.78 3.22a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0ZM22.5 9.752a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75ZM14.249 26a4.747 4.747 0 0 1-4.068-2.296l-2.01 1.006a2.75 2.75 0 0 1-3.174-.516l-1.19-1.19a2.75 2.75 0 0 1-.516-3.174l7.155-14.31a2.75 2.75 0 0 1 4.404-.714l8.345 8.345a2.75 2.75 0 0 1-.715 4.404l-3.827 1.913A4.75 4.75 0 0 1 14.25 26Zm3.056-5.857l-5.776 2.887a3.25 3.25 0 0 0 5.776-2.888Z"/>`),
		g.Group(children),
	)
}