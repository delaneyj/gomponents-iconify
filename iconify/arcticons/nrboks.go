package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nrboks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="21" height="21" x="13.5" y="13.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/>`),
		g.Group(children),
	)
}