package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Riseup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.5 24l-1.485-9.007l-8.961-1.738L8.16 9.06L7.045 0l6.495 6.414l8.271-3.861l-4.093 8.16l6.228 6.673l-9.024-1.372z"/>`),
		g.Group(children),
	)
}