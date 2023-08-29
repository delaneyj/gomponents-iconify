package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.374 10.168a.5.5 0 0 0-.748.664l4 4.5a.5.5 0 0 0 .728.022l10.5-10.5a.5.5 0 0 0-.707-.708L7.02 14.271l-3.647-4.103Z"/>`),
		g.Group(children),
	)
}