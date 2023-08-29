package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListAlphabet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 9h22M20 19h22M20 29h22M20 39h22M6 29h6L6 39h6"/><path d="M11 9H7l-.7 7h5.4L11 9Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 19l.3-3m5.7 3l-.3-3m0 0L11 9H7l-.7 7m5.4 0H6.3"/></g>`),
		g.Group(children),
	)
}