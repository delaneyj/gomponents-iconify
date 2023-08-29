package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miremote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="7.981" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.816" cy="24" r=".75" fill="currentColor"/><circle cx="24" cy="9.185" r=".75" fill="currentColor"/><circle cx="9.184" cy="24" r=".75" fill="currentColor"/><circle cx="24" cy="38.816" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}