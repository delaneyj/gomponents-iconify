package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneBooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 34h20v9H14zm0-30h20v6H14zm0 6v24m6-24v24m2-18v4m10 6H14m20-16v24M4 44h40"/>`),
		g.Group(children),
	)
}